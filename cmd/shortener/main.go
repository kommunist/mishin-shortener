package main

import (
	"log"

	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/handlers"
	middleware "mishin-shortener/internal/app/midleware"
	"mishin-shortener/internal/app/storage"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	c := config.MakeConfig()
	c.InitConfig()
	db := storage.MakeDatabase()
	h := handlers.MakeShortanerHandler(&c, &db)

	r := chi.NewRouter()
	r.Use(middleware.WithLogRequest)

	r.Post("/", h.CreateURLHandler)
	r.Post("/api/shorten", h.CreateURLByJSONHandler)
	r.Get("/{shortened}", h.RedirectHandler)

	log.Printf("Server started on %s", c.BaseServerURL)
	err := http.ListenAndServe(c.BaseServerURL, r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
