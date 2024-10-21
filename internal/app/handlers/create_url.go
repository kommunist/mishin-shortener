package handlers

import (
	"io"
	"log/slog"
	"mishin-shortener/internal/app/exsist"
	"mishin-shortener/internal/app/hasher"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

// Обработчик простого запроса на сокращение.
func (h *ShortanerHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	status := http.StatusCreated
	hashed := hasher.GetMD5Hash(body)

	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	err = h.DB.Push(r.Context(), hashed, string(body), userID)
	if err != nil {
		if _, ok := err.(*exsist.ExistError); ok { // обрабатываем проблему, когда уже есть в базе
			status = http.StatusConflict
		} else {
			slog.Error("push to storage error", "err", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
	}

	w.WriteHeader(status)
	w.Write(h.resultURL(hashed))
}
