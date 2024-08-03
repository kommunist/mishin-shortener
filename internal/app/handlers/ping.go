package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

func (h *ShortanerHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := h.Driver.PingContext(ctx); err != nil {
		slog.Error("error when ping database", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
