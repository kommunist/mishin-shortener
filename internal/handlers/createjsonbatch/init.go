package createjsonbatch

import (
	"context"
	"mishin-shortener/internal/config"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type Pusher interface {
	PushBatch(context.Context, *map[string]string, string) error // collection, userID
}

// Структура хендлера
type Handler struct {
	storage  Pusher
	settings config.MainConfig

	pb.UnimplementedCreateBatchServer
}

// Конструктор хендлера
func Make(settings config.MainConfig, storage Pusher) Handler {
	return Handler{storage: storage, settings: settings}
}

type requestBatchItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type responseBatchItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
