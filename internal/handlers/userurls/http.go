package userurls

import (
	"encoding/json"
	"log/slog"
	"mishin-shortener/internal/secure"
	"net/http"
)

// Обработчик, возвращающий все сокращенные урлы пользователя.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	data, err := h.Perform(r.Context(), userID)
	if err != nil {
		http.Error(w, "Error when get data from db", http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error when create json", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		slog.Error("error when write response", "err", err)
		http.Error(w, "Write response error", http.StatusInternalServerError)
		return
	}
}
