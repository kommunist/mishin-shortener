package handlers

import (
	"io"
	"mishin-shortener/internal/app/hasher"
	"net/http"
)

func (h *ShortanerHandler) CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashed := hasher.GetMD5Hash(body)

	h.Db.Push("/"+hashed, string(body))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(h.Options.BaseRedirectURL + "/" + hashed))
}
