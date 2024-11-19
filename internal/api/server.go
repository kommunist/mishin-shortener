package api

import (
	"context"
	"log/slog"
	"net/http"
)

// Основной метод пакета API
func (a *ShortanerAPI) Call() error {
	a.initServ()

	slog.Info("server started", "URL", a.setting.BaseServerURL)

	if a.setting.EnableHTTPS {
		return a.startWithTLS()
	} else {
		return a.start()
	}
}

func (a *ShortanerAPI) start() error {
	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "err", err)
		return err
	}
	return nil
}

func (a *ShortanerAPI) startWithTLS() error {

	err := a.Server.ListenAndServeTLS("certs/MyCertificate.crt", "certs/MyKey.key")
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start with tls", "err", err)
		return err
	}
	return nil
}

// Функция, останавливающая сервер
func (a *ShortanerAPI) Stop() {
	err := a.Server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error when shutdown server", "err", err)
	}

	err = a.storage.Finish()
	if err != nil { // будем счиать, что отвественность api закрыть базу при выключении
		slog.Error("Error when close connection to storage", "err", err)
	}
}
