package handlers

import (
	"io"
	"log/slog"
	"mishin-shortener/internal/app/hasher"
	"net/http"
)

func (h *ShortanerHandler) CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashed := hasher.GetMD5Hash(body)

	err = h.DB.Push("/"+hashed, string(body))
	if err != nil {
		slog.Error("push to storage error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(h.Options.BaseRedirectURL + "/" + hashed))
}
