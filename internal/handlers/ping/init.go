package ping

import (
	"context"
)

// Интерфейс доступа к базе
type Pinger interface {
	Ping(context.Context) error
}

// Структура хендлера
type Handler struct {
	storage Pinger
}

// Конструктор хендлера
func Make(storage Pinger) Handler {
	return Handler{storage: storage}
}
