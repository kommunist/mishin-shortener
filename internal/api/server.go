package api

import (
	"context"
	"log/slog"
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
	if err != nil {
		slog.Error("Server failed to start", "err", err)
		return err
	}
	return nil
}

func (a *ShortanerAPI) startWithTLS() error {

	err := a.Server.ListenAndServeTLS("certs/MyCertificate.crt", "certs/MyKey.key")
	if err != nil {
		slog.Error("Server failed to start with tls", "err", err)
		return err
	}
	return nil
}

func (a *ShortanerAPI) Stop() {
	err := a.Server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error when shutdown server", "err", err)
	}
}
