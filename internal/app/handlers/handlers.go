package handlers

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/storage"
)

type ShortanerHandler struct {
	DB      *storage.Database
	Options *config.MainConfig
}

func MakeShortanerHandler(c *config.MainConfig, db *storage.Database) ShortanerHandler {
	return ShortanerHandler{
		DB:      db,
		Options: c,
	}
}
