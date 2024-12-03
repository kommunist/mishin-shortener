package redirect

import (
	"context"
	pb "mishin-shortener/proto"
)

// Интерфейс доступа к базе
type Getter interface {
	Get(context.Context, string) (string, error)
}

// Структура хендлера
type Handler struct {
	storage Getter

	pb.UnimplementedGetServer
}

// Конструктор хендлера
func Make(storage Getter) Handler {
	return Handler{storage: storage}
}
