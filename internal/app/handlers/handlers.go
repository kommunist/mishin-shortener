package handlers

import (
	"context"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/delasync"
)

// Интерфейс, описывающий работу с базой.
type AbstractStorage interface {
	Push(context.Context, string, string, string) error          // убрать и перегенерировать моки. Пока оставлено ради тестов
	PushBatch(context.Context, *map[string]string, string) error // collection, userID
	Get(context.Context, string) (string, error)
	GetByUserID(context.Context, string) (map[string]string, error) // userID
	DeleteByUserID(context.Context, []delasync.DelPair) error       // слайс пар userID, list
	Finish() error
	Ping(context.Context) error
}

// Структура обработчика, связывающая базу и конфигурацию.
type ShortanerHandler struct {
	DB      AbstractStorage
	Options config.MainConfig
}

// Создание структуры обработчика.
func MakeShortanerHandler(c config.MainConfig, db AbstractStorage) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
	}
}

func (h *ShortanerHandler) resultURL(hashed string) []byte {
	return []byte(h.Options.BaseRedirectURL + "/" + hashed)
}
