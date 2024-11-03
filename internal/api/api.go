package api

import (
	"log/slog"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/delasync"
	"mishin-shortener/internal/app/handlers"
	middleware "mishin-shortener/internal/app/midleware"
	"mishin-shortener/internal/handlers/userurls"
	"net/http"
	"os"
	"time"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
)

// Основная структуруа пакета API
type ShortanerAPI struct {
	setting config.MainConfig
	storage handlers.AbstractStorage // пока используем общий интерфейс. Потом сделаем композицию

	userUrls userurls.Handler
}

// Конструктор структуры пакета API
func Make(setting config.MainConfig, storage handlers.AbstractStorage) ShortanerAPI {
	return ShortanerAPI{
		setting: setting,
		storage: storage,

		userUrls: userurls.Make(setting, storage),
	}
}

// Основной метод пакета API
func (a *ShortanerAPI) Call() {
	h := handlers.MakeShortanerHandler(a.setting, a.storage)

	delasync.InitWorker(h.DelChan, h.DB.DeleteByUserID) // не дело из api запускать асинхрон. Но пока так

	r := chi.NewRouter()

	r.Use(chiMiddleware.Timeout(60 * time.Second))
	r.Use(middleware.WithLogRequest)
	r.Use(middleware.Gzip)

	r.Route("/api", func(r chi.Router) {
		r.With(middleware.AuthSet).Route("/shorten", func(r chi.Router) {
			r.Post("/", h.CreateURLByJSON)
			r.Post("/batch", h.CreateURLByJSONBatch)
		})

		r.With(middleware.AuthCheck).Route("/user", func(r chi.Router) {
			r.Get("/urls", a.userUrls.Call)
			r.Delete("/urls", h.DeleteURLs)
		})

	})
	r.With(middleware.AuthSet).Post("/", h.CreateURL)
	r.Get("/{shortened}", h.RedirectHandler)
	r.Get("/ping", h.PingHandler)

	slog.Info("server started", "URL", a.setting.BaseServerURL)

	err := http.ListenAndServe(a.setting.BaseServerURL, r)
	if err != nil {
		slog.Error("Server failed to start", "err", err)
		os.Exit(1)
	}

}
