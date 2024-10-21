package handlers

import (
	"log/slog"
	"net/http"
)

// Обработчик для проверки и восстановления подключения к базе.
func (h *ShortanerHandler) PingHandler(w http.ResponseWriter, r *http.Request) {

	if h.Options.DatabaseDSN == "" {
		slog.Error("database DSN not defined for ping")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	if err := h.DB.Ping(r.Context()); err != nil {
		slog.Error("error when ping database", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
