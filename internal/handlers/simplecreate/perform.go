package simplecreate

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/hasher"
)

func (h *Handler) Perform(ctx context.Context, original []byte, userID string) (short string, err error) {
	short = hasher.GetMD5Hash(original)

	err = h.storage.Push(ctx, short, string(original), userID)
	if err != nil {
		slog.Error("Error when push data to storage")
		return short, err
	}

	return short, nil
}
