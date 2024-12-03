package ping

import (
	"context"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type Pinger interface {
	Ping(context.Context) error
}

// Структура хендлера
type Handler struct {
	storage Pinger

	pb.UnimplementedPingServer
}

// Конструктор хендлера
func Make(storage Pinger) Handler {
	return Handler{storage: storage}
}
