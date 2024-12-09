package ping

import (
	"context"
	"log/slog"
)

func (h *Handler) Perform(ctx context.Context) error {
	err := h.storage.Ping(ctx)
	if err != nil {
		slog.Error("Error when ping storage")
		return err
	}
	return nil
}
