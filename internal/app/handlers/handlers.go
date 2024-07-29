package handlers

import (
	"mishin-shortener/internal/app/config"
)

type AbstractStorage interface {
	Push(string, string) error
	Get(string) (string, error)
	Finish() error
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
