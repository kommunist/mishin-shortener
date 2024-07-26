package handlers

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/storage"
)

type ShortanerHandler struct {
	DB      storage.Abstract
	Options config.MainConfig
}

func MakeShortanerHandler(c config.MainConfig, db storage.Abstract) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
	}
}
