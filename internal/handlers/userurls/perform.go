package userurls

import (
	"context"
	"log/slog"
)

func (h *Handler) Perform(ctx context.Context, userID string) ([]responseItem, error) {
	data, err := h.storage.GetByUserID(ctx, userID)
	if err != nil {
		slog.Error("Error when get data from db", "err", err)
		return nil, err
	}

	result := make([]responseItem, 0)
	for k, v := range data {
		result = append(
			result,
			responseItem{Short: h.setting.BaseRedirectURL + "/" + k, Original: v},
		)
	}

	return result, nil
}
