package stats

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseItem struct {
	Users int `json:"users"`
	Urls  int `json:"urls"`
}

func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	realIP := r.Header.Get("X-Real-IP")
	if !h.netChecker.Contains(realIP) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	users, urls, err := h.storage.GetStats(r.Context())
	if err != nil {
		slog.Error("Error when get data from storage", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	item := ResponseItem{Users: users, Urls: urls}

	data, err := json.Marshal(item)
	if err != nil {
		slog.Error("Error when serialize json", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		slog.Error("Error when write data to response", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
