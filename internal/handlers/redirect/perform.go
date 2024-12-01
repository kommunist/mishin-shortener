package redirect

import (
	"context"
	"log/slog"
)

func (h *Handler) Perform(ctx context.Context, short string) (string, error) {
	to, err := h.storage.Get(ctx, short)

	if err != nil {
		slog.Error("Error when get data from storage", "err", err)
		return "", err
	}

	return to, nil
}
