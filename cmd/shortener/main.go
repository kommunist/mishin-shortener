package main

import (
	"context"
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

	go func(in <-chan [2]string) {
		var buf [][2]string // сюда будем складывать накопленные

		rf := func(in <-chan [2]string) ([2]string, bool) {
			select {
			case val := <-in:
				return val, true
			case <-time.After(5 * time.Second):
				return [2]string{}, false
			}
		}

		for {
			val, found := rf(in)
			if found {
				buf = append(buf, val)
				if len(buf) > 2 {
					h.DB.DeleteByUserID(context.Background(), buf)
					buf = nil
				}
			} else {
				if len(buf) > 0 {
					h.DB.DeleteByUserID(context.Background(), buf)
					buf = nil
				}
			}
		}
	}(h.DelChan)

	r.Use(chiMiddleware.Timeout(60 * time.Second))
	r.Use(middleware.WithLogRequest)
	r.Use(middleware.GzipMiddleware)

	r.Route("/api", func(r chi.Router) {
		r.With(middleware.AuthSet).Route("/shorten", func(r chi.Router) {
			r.Post("/", h.CreateURLByJSON)
			r.Post("/batch", h.CreateURLByJSONBatch)
		})

		r.With(middleware.AuthCheck).Route("/user", func(r chi.Router) {
			r.Get("/urls", h.UserURLs)
			r.Delete("/urls", h.DeleteURLs)
		})

	})
	r.With(middleware.AuthSet).Post("/", h.CreateURL)
	r.Get("/{shortened}", h.RedirectHandler)
	r.Get("/ping", h.PingHandler)

	slog.Info("server started", "URL", c.BaseServerURL)

	err := http.ListenAndServe(c.BaseServerURL, r)
	if err != nil {
		slog.Error("Server failed to start", "err", err)
		os.Exit(1)
	}
}
