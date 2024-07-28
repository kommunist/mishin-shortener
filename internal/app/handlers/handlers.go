package handlers

import (
	"mishin-shortener/internal/app/config"
)

type abstractStorage interface {
	Push(string, string) error
	Get(string) (string, error)
}

type ShortanerHandler struct {
	DB      abstractStorage
	Options config.MainConfig
}

func MakeShortanerHandler(c config.MainConfig, db abstractStorage) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
	}
}
