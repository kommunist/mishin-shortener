package api

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/delasync"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/handlers/deleteurls"
	"mishin-shortener/internal/handlers/simplecreate"
	"mishin-shortener/internal/handlers/userurls"
	"net/http"
)

// Основная структуруа пакета API
type ShortanerAPI struct {
	setting config.MainConfig
	storage handlers.AbstractStorage // пока используем общий интерфейс. Потом сделаем композицию

	delChan chan delasync.DelPair // [0] - для user_id и [1] для short

	userUrls     userurls.Handler
	simpleCreate simplecreate.Handler
	deleteURLs   deleteurls.Handler

	Server http.Server
}

// Конструктор структуры пакета API
func Make(setting config.MainConfig, storage handlers.AbstractStorage) ShortanerAPI {
	c := make(chan delasync.DelPair, 5)
	return ShortanerAPI{
		setting: setting,
		storage: storage,

		delChan: c, // по большому счету оно тут для того, чтобы закрыть

		userUrls:     userurls.Make(setting, storage),
		simpleCreate: simplecreate.Make(setting, storage),
		deleteURLs:   deleteurls.Make(c),
	}
}

func (a *ShortanerAPI) initServ() {
	router := a.initRouter()
	a.Server = http.Server{
		Addr:    a.setting.BaseServerURL,
		Handler: router,
	}
}
