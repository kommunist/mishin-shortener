package simplecreate

import (
	"context"
	"mishin-shortener/internal/config"
)

// Интерфейс доступа к базе
type Pusher interface {
	Push(context.Context, string, string, string) error
}

// Структура хендлера
type Handler struct {
	storage Pusher
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage Pusher) Handler {
	return Handler{storage: storage, setting: setting}
}
