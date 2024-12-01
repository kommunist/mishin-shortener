package deleteurls

import (
	"encoding/json"
	"io"
	"log/slog"
	"mishin-shortener/internal/secure"
	"net/http"
)

// Обработчик запроса на удаление сокращенного URL.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := make([]string, 0)
	err = json.Unmarshal(body, &list)
	if err != nil {
		slog.Error("Error while parsing json")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	h.Perform(r.Context(), list, userID)

	w.WriteHeader(http.StatusAccepted)
}
