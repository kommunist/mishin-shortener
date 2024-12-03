package httpserver

import (
	middleware "mishin-shortener/internal/midleware"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
)

func (a *HTTPHandler) initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Timeout(60 * time.Second))
	r.Use(middleware.WithLogRequest)
	r.Use(middleware.Gzip)

	r.Route("/api", func(r chi.Router) {
		r.With(middleware.AuthSet).Route("/shorten", func(r chi.Router) {
			r.Post("/", a.createJSON.Call)
			r.Post("/batch", a.createJSONBatch.Call)
		})

		r.With(middleware.AuthCheck).Route("/user", func(r chi.Router) {
			r.Get("/urls", a.userUrls.Call)
			r.Delete("/urls", a.deleteURLs.Call)
		})

	})
	r.With(middleware.AuthSet).Post("/", a.simpleCreate.Call)
	r.Get("/{shortened}", a.redirect.Call)
	r.Get("/ping", a.ping.Call)

	return r
}
