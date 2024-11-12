package userurls

import (
	"encoding/json"
	"log/slog"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

type responseItem struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

// Обработчик, возвращающий все сокращенные урлы пользователя.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(secure.UserIDKey)
	if u == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userID := u.(string)

	data, err := h.storage.GetByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Error when get data from db", http.StatusInternalServerError)
		return
	}
	result := make([]responseItem, 0)

	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	for k, v := range data {
		result = append(
			result,
			responseItem{Short: h.setting.BaseRedirectURL + "/" + k, Original: v},
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	out, err := json.Marshal(result)
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
