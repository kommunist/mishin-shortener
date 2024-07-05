package main

import (
	"net/http"

	"internal/handlers"
	"internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {

	db := storage.Database{}
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, &db)
	})
	r.Get("/{shortened}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHandler(w, r, &db)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
