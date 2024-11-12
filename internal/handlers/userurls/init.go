package userurls

import (
	"context"
	"mishin-shortener/internal/app/config"
)

// Интерфейс доступа к базе
type ByUserIDGetter interface {
	GetByUserID(context.Context, string) (map[string]string, error)
}

// Структура хендлера
type Handler struct {
	storage ByUserIDGetter
	setting config.MainConfig
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage ByUserIDGetter) Handler {
	return Handler{storage: storage, setting: setting}
}
