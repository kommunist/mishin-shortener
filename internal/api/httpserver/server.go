package httpserver

import (
	"context"
	"log/slog"
	"net/http"
)

// Основной метод пакета API
func (a *HTTPHandler) Call() error {
	slog.Info("server started", "URL", a.setting.BaseServerURL)

	if a.setting.EnableHTTPS {
		return a.startWithTLS()
	} else {
		return a.start()
	}
}

func (a *HTTPHandler) start() error {
	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "err", err)
		return err
	}
	return nil
}

func (a *HTTPHandler) startWithTLS() error {
	err := a.Server.ListenAndServeTLS("/etc/ssl/self_created/MyCertificate.crt", "/etc/ssl/self_created/MyKey.key")
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start with tls", "err", err)
		return err
	}
	return nil
}

// Функция, останавливающая сервер
func (a *HTTPHandler) Stop() {
	err := a.Server.Shutdown(context.Background())
	if err != nil {
		slog.Error("Error when shutdown server", "err", err)
	}
}
