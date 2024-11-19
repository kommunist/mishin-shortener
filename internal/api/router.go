package api

import (
	"mishin-shortener/internal/app/delasync"
	"mishin-shortener/internal/app/handlers"
	middleware "mishin-shortener/internal/app/midleware"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (a *ShortanerAPI) initRouter() *chi.Mux {
	h := handlers.MakeShortanerHandler(a.setting, a.storage)

	delasync.InitWorker(a.delChan, h.DB.DeleteByUserID) // не дело из api запускать асинхрон. Но пока так

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
			r.Delete("/urls", a.deleteURLs.Call)
		})

	})
	r.With(middleware.AuthSet).Post("/", a.simpleCreate.Call)
	r.Get("/{shortened}", h.RedirectHandler)
	r.Get("/ping", h.PingHandler)

	return r
}
