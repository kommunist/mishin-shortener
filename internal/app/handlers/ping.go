package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func (h *ShortanerHandler) PingHandler(w http.ResponseWriter, r *http.Request) {

	if h.Options.DatabaseDSN == "" {
		slog.Error("database DSN not defined for ping")
		// поменять на другой тип ошибки
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := h.DB.Ping(ctx); err != nil {
		slog.Error("error when ping database", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
