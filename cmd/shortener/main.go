package main

import (
	"net/http"

	"internal/config"
	"internal/handlers"
	"internal/storage"

	"fmt"
	"github.com/go-chi/chi/v5"
)

func main() {

	db := storage.Database{}
	r := chi.NewRouter()

	config.Parse()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.PostHandler(w, r, &db, config.Config.BaseRedirectUrl)
	})
	r.Get("/{shortened}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHandler(w, r, &db)
	})

	fmt.Println("Server started on", config.Config.BaseServerUrl)
	err := http.ListenAndServe(config.Config.BaseServerUrl, r)
	if err != nil {
		panic(err)
	}
}
