package handlers

import (
	"mishin-shortener/internal/app/storage"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request, db *storage.Database) {
	toLocation := db.Get(r.RequestURI)

	if toLocation != "" {
		w.Header().Set("Location", toLocation)
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
