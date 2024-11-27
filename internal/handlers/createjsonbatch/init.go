package createjsonbatch

import (
	"context"
	"mishin-shortener/internal/config"
)

// Интерфейс доступа к базе
type Pusher interface {
	PushBatch(context.Context, *map[string]string, string) error // collection, userID
}

// Структура хендлера
type Handler struct {
	storage  Pusher
	settings config.MainConfig
}

// Конструктор хендлера
func Make(settings config.MainConfig, storage Pusher) Handler {
	return Handler{storage: storage, settings: settings}
}
