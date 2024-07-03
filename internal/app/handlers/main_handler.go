package handlers

import (
	"fmt"
	"internal/hasher"
	"internal/storage"
	"io"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request, db *storage.Database) {
	if r.Method == http.MethodPost && r.RequestURI == "/" {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		hashed := hasher.GetMD5Hash(body)

		db.Push("/"+hashed, string(body))

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("http://" + r.Host + "/" + hashed))

	} else if toLocation := db.Get(r.RequestURI); r.Method == http.MethodGet && toLocation != "" {
		fmt.Println(toLocation)
		w.Header().Set("Location", toLocation)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
