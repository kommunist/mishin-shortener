package handlers

import (
	"context"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/delasync"
)

// Интерфейс, описывающий работу с базой.
type AbstractStorage interface {
	Push(context.Context, string, string, string) error          // short, original, userID
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
	DelChan chan delasync.DelPair // [0] - для user_id и [1] для short
}

// Создание структуры обработчика.
func MakeShortanerHandler(c config.MainConfig, db AbstractStorage) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
		DelChan: make(chan delasync.DelPair, 5),
	}
}

func (h *ShortanerHandler) resultURL(hashed string) []byte {
	return []byte(h.Options.BaseRedirectURL + "/" + hashed)
}
