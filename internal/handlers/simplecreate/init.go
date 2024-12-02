package simplecreate

import (
	"context"
	"mishin-shortener/internal/config"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type Pusher interface {
	Push(context.Context, string, string, string) error
}

// Структура хендлера
type Handler struct {
	storage Pusher
	setting config.MainConfig

	pb.UnimplementedCreateServer
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage Pusher) Handler {
	return Handler{storage: storage, setting: setting}
}
