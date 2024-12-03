package deleteurls

import (
	"context"
	"mishin-shortener/internal/delasync"
)

func (h *Handler) Perform(ctx context.Context, list []string, userID string) {
	for _, v := range list {
		h.DelChan <- delasync.DelPair{UserID: userID, Item: v}
	}
}
