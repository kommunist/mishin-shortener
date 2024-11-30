package stats

import (
	"context"
	"log/slog"
)

// Основной метод с бизнесс-логикой приложения
func (h *Handler) Perform(ctx context.Context) (users int, urls int, err error) {
	users, urls, err = h.storage.GetStats(ctx)
	if err != nil {
		slog.Error("Error when get data from storage", "err", err)
		return 0, 0, err
	}
	return users, urls, nil
}
