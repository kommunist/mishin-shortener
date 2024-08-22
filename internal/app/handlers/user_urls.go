package handlers

import (
	"encoding/json"
	"log/slog"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

type UserURLsItem struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func (h *ShortanerHandler) UserURLs(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(secure.UserIDKey)
	slog.Info("User id in context", "user_id", u)
	if u == nil {
		return
	}
	userID := u.(string)

	data, err := h.DB.GetByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Error when get data from db", http.StatusInternalServerError)
		return
	}
	result := make([]UserURLsItem, 0)

	if len(data) > 0 {
		for k, v := range data {
			result = append(result, UserURLsItem{Short: h.Options.BaseRedirectURL + k, Original: v})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		out, err := json.Marshal(result)
		if err != nil {
			http.Error(w, "Error when create json", http.StatusInternalServerError)
			return
		}

		w.Write(out)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
