package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"mishin-shortener/internal/app/delasync"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

// Обработчик запроса на удаление сокращенного URL.
func (h *ShortanerHandler) DeleteURLs(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(secure.UserIDKey)
	slog.Info("User id in context", "user_id", u)
	if u == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userID := u.(string)

	list := make([]string, 0)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &list)
	if err != nil {
		slog.Error("Error while parsing json")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	for _, v := range list {
		h.DelChan <- delasync.DelPair{UserID: userID, Item: v}
	}

	w.WriteHeader(http.StatusAccepted)
}
