package grpcserver

import (
	"log/slog"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/handlers/createjson"
	"mishin-shortener/internal/handlers/createjsonbatch"
	"mishin-shortener/internal/handlers/deleteurls"
	"mishin-shortener/internal/handlers/ping"
	"mishin-shortener/internal/handlers/redirect"
	"mishin-shortener/internal/handlers/simplecreate"
	"mishin-shortener/internal/handlers/stats"
	"mishin-shortener/internal/handlers/userurls"
	"net"

	pb "mishin-shortener/proto"

	"google.golang.org/grpc"
)

// Композиция интерфейсов для доступа в базу
type CommonStorage interface {
	userurls.ByUserIDGetter
	simplecreate.Pusher
	redirect.Getter
	ping.Pinger
	createjson.Pusher
	createjsonbatch.Pusher
	stats.StatsGetter
}

// Основная структуруа пакета API
type GRPCHandler struct {
	setting config.MainConfig

	userUrls        userurls.Handler
	simpleCreate    simplecreate.Handler
	deleteURLs      deleteurls.Handler
	redirect        redirect.Handler
	ping            ping.Handler
	createJSON      createjson.Handler
	createJSONBatch createjsonbatch.Handler
	stats           stats.Handler

	listener net.Listener
	server   *grpc.Server

	pb.UnimplementedShortenerServer // добавим сюда же работу grpc
}

// Конструктор структуры пакета GRPC
func Make(setting config.MainConfig, storage CommonStorage, c chan delasync.DelPair) *GRPCHandler {
	h := GRPCHandler{
		setting:         setting,
		userUrls:        userurls.Make(setting, storage),
		simpleCreate:    simplecreate.Make(setting, storage),
		deleteURLs:      deleteurls.Make(c),
		redirect:        redirect.Make(storage),
		ping:            ping.Make(storage),
		createJSON:      createjson.Make(setting, storage),
		createJSONBatch: createjsonbatch.Make(setting, storage),
	}

	stats, err := stats.Make(setting, storage)
	if err != nil {
		slog.Error("Error when make stats handler", "err", err)
		panic(err) // сделать вынос ошибки
	}
	h.stats = stats

	listener, err := net.Listen("tcp", ":3200")
	if err != nil {
		slog.Error("Error when listen net", "err", err) // сделать возврат err
	}
	h.listener = listener

	h.server = grpc.NewServer()

	pb.RegisterShortenerServer(h.server, &h)

	return &h
}
