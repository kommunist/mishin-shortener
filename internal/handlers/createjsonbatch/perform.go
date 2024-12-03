package createjsonbatch

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/hasher"
)

func (h *Handler) Perform(ctx context.Context, list []requestBatchItem, userID string) ([]responseBatchItem, error) {
	prepareToSave := make(map[string]string)          // это для отдачи в базу
	output := make([]responseBatchItem, 0, len(list)) // а это для результата

	for _, v := range list {
		hashed := hasher.GetMD5Hash([]byte(v.OriginalURL))

		prepareToSave[hashed] = v.OriginalURL

		output = append(
			output,
			responseBatchItem{
				CorrelationID: v.CorrelationID,
				ShortURL:      h.settings.BaseRedirectURL + "/" + hashed,
			},
		)
	}

	err := h.storage.PushBatch(ctx, &prepareToSave, userID)
	if err != nil {
		slog.Error("Error when push data to storage", "err", err)
		return nil, err
	}
	return output, nil
}
