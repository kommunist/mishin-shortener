package createjson

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
	setting config.MainConfig
	storage Pusher
}

// Конструктор хендлера
func Make(setting config.MainConfig, storage Pusher) Handler {
	return Handler{storage: storage, setting: setting}
}

func (h *Handler) resultURL(hashed string) []byte {
	return []byte(h.setting.BaseRedirectURL + "/" + hashed)
}
