package handlers

import (
	"context"
	"mishin-shortener/internal/app/config"
)

type AbstractStorage interface {
	Push(context.Context, string, string, string) error          // short, original, userId
	PushBatch(context.Context, *map[string]string, string) error // collection, userId
	Get(context.Context, string) (string, error)
	Finish() error
	Ping(context.Context) error
}

type ShortanerHandler struct {
	DB      AbstractStorage
	Options config.MainConfig
}

func MakeShortanerHandler(c config.MainConfig, db AbstractStorage) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
	}
}

func (h *ShortanerHandler) resultURL(hashed string) []byte {
	return []byte(h.Options.BaseRedirectURL + "/" + hashed)
}
