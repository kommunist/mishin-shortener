package api

import (
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/handlers/createjson"
	"mishin-shortener/internal/handlers/createjsonbatch"
	"mishin-shortener/internal/handlers/deleteurls"
	"mishin-shortener/internal/handlers/ping"
	"mishin-shortener/internal/handlers/redirect"
	"mishin-shortener/internal/handlers/simplecreate"
	"mishin-shortener/internal/handlers/userurls"
	"net/http"
)

// Композиция интерфейсов для доступа в базу
type CommonStorage interface {
	userurls.ByUserIDGetter
	simplecreate.Pusher
	redirect.Getter
	ping.Pinger
	createjson.Pusher
	createjsonbatch.Pusher
}

// Основная структуруа пакета API
type ShortanerAPI struct {
	setting config.MainConfig

	userUrls        userurls.Handler
	simpleCreate    simplecreate.Handler
	deleteURLs      deleteurls.Handler
	redirect        redirect.Handler
	ping            ping.Handler
	createJSON      createjson.Handler
	createJSONBatch createjsonbatch.Handler

	Server http.Server
}

// Конструктор структуры пакета API
func Make(setting config.MainConfig, storage CommonStorage, c chan delasync.DelPair) *ShortanerAPI {
	api := ShortanerAPI{
		setting:         setting,
		userUrls:        userurls.Make(setting, storage),
		simpleCreate:    simplecreate.Make(setting, storage),
		deleteURLs:      deleteurls.Make(c),
		redirect:        redirect.Make(storage),
		ping:            ping.Make(storage),
		createJSON:      createjson.Make(setting, storage),
		createJSONBatch: createjsonbatch.Make(setting, storage),
	}

	api.Server = http.Server{
		Addr:    setting.BaseServerURL,
		Handler: api.initRouter(),
	}

	return &api
}
