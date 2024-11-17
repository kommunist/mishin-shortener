package api

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/handlers/simplecreate"
	"mishin-shortener/internal/handlers/userurls"
	"net/http"
	"os"
)

// Основная структуруа пакета API
type ShortanerAPI struct {
	setting config.MainConfig
	storage handlers.AbstractStorage // пока используем общий интерфейс. Потом сделаем композицию

	userUrls     userurls.Handler
	simpleCreate simplecreate.Handler

	server  http.Server
	intChan chan os.Signal
}

// Конструктор структуры пакета API
func Make(setting config.MainConfig, storage handlers.AbstractStorage, c chan os.Signal) ShortanerAPI {
	return ShortanerAPI{
		setting: setting,
		storage: storage,

		userUrls:     userurls.Make(setting, storage),
		simpleCreate: simplecreate.Make(setting, storage),
		intChan:      c,
	}
}

func (a *ShortanerAPI) initServ() {
	router := a.initRouter()
	a.server = http.Server{
		Addr:    a.setting.BaseServerURL,
		Handler: router,
	}
}
