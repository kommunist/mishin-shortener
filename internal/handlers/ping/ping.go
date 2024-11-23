package ping

import (
	"log/slog"
	"net/http"
)

// Обработчик для проверки и восстановления подключения к базе.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	err := h.storage.Ping(r.Context())
	if err != nil {
		slog.Error("error when ping database", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
