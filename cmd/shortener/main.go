package main

import (
	"log/slog"
	"os"

	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/filestorage"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/app/mapstorage"
	middleware "mishin-shortener/internal/app/midleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func setStorage(c config.MainConfig) handlers.AbstractStorage {
	if c.FileStoragePath != "" {
		return filestorage.Make(c.FileStoragePath)
	} else {
		return mapstorage.Make()
	}
}

func main() {
	c := config.MakeConfig()
	c.InitConfig()

	storage := setStorage(c)

	h := handlers.MakeShortanerHandler(c, storage)

	r := chi.NewRouter()
	r.Use(middleware.WithLogRequest)
	r.Use(middleware.GzipMiddleware)

	r.Post("/", h.CreateURLHandler)
	r.Post("/api/shorten", h.CreateURLByJSONHandler)
	r.Get("/{shortened}", h.RedirectHandler)

	slog.Info("server started", "URL", c.BaseServerURL)

	err := http.ListenAndServe(c.BaseServerURL, r)
	if err != nil {
		slog.Error("Server failed to start", "err", err)
		os.Exit(1)
	}
}
