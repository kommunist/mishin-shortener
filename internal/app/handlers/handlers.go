package handlers

import (
	"context"
	"mishin-shortener/internal/app/config"
)

type AbstractStorage interface {
	Push(string, string) error
	PushBatch(*map[string]string) error
	Get(string) (string, error)
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

func (h *ShortanerHandler) resultUrl(hashed string) []byte {
	return []byte(h.Options.BaseRedirectURL + "/" + hashed)
}
