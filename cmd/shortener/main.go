package main

import (
	"log/slog"
	"os"
	"time"

	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/filestorage"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/app/mapstorage"
	middleware "mishin-shortener/internal/app/midleware"
	"mishin-shortener/internal/app/pgstorage"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func initStorage(c config.MainConfig) handlers.AbstractStorage {
	if c.DatabaseDSN != "" {
		return pgstorage.Make(c)
	}
	if c.FileStoragePath != "" {
		return filestorage.Make(c.FileStoragePath)
	}

	return mapstorage.Make()
}

func main() {
	c := config.MakeConfig()
	c.InitConfig()

	storage := initStorage(c)
	defer storage.Finish()

	h := handlers.MakeShortanerHandler(c, storage)

	r := chi.NewRouter()
	r.Use(chiMiddleware.Timeout(60 * time.Second))
	r.Use(middleware.WithLogRequest)
	r.Use(middleware.GzipMiddleware)
	r.Use(middleware.AuthMiddleware)

	r.Post("/", h.CreateURL)
	r.Post("/api/shorten", h.CreateURLByJSON)
	r.Post("/api/shorten/batch", h.CreateURLByJSONBatch)
	r.Get("/{shortened}", h.RedirectHandler)
	r.Get("/ping", h.PingHandler)
	r.Get("/api/user/urls", h.UserURLs)

	slog.Info("server started", "URL", c.BaseServerURL)

	err := http.ListenAndServe(c.BaseServerURL, r)
	if err != nil {
		slog.Error("Server failed to start", "err", err)
		os.Exit(1)
	}
}
