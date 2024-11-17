package api

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/handlers/simplecreate"
	"mishin-shortener/internal/handlers/userurls"
	"net/http"
)

// Основная структуруа пакета API
type ShortanerAPI struct {
	setting config.MainConfig
	storage handlers.AbstractStorage // пока используем общий интерфейс. Потом сделаем композицию

	userUrls     userurls.Handler
	simpleCreate simplecreate.Handler

	Server http.Server
}

// Конструктор структуры пакета API
func Make(setting config.MainConfig, storage handlers.AbstractStorage) ShortanerAPI {
	return ShortanerAPI{
		setting: setting,
		storage: storage,

		userUrls:     userurls.Make(setting, storage),
		simpleCreate: simplecreate.Make(setting, storage),
	}
}

func (a *ShortanerAPI) initServ() {
	router := a.initRouter()
	a.Server = http.Server{
		Addr:    a.setting.BaseServerURL,
		Handler: router,
	}
}
