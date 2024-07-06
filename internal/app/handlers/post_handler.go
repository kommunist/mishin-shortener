package handlers

import (
	"internal/hasher"
	"internal/storage"
	"io"
	"net/http"
)

func PostHandler(w http.ResponseWriter, r *http.Request, db *storage.Database, redirectHost string) {
	if r.RequestURI == "/" {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		hashed := hasher.GetMD5Hash(body)

		db.Push("/"+hashed, string(body))

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(redirectHost + "/" + hashed))

	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
