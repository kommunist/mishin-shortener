package redirect

import (
	"context"
)

// Интерфейс доступа к базе
type Getter interface {
	Get(context.Context, string) (string, error)
}

// Структура хендлера
type Handler struct {
	storage Getter
}

// Конструктор хендлера
func Make(storage Getter) Handler {
	return Handler{storage: storage}
}
