package main

import (
	"net/http"

	"internal/handlers"
	"internal/storage"
)

func main() {

	db := storage.Database{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.MainHandler(w, r, &db)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
