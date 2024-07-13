package handlers

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/storage"
)

type ShortanerHandler struct {
	Db      *storage.Database
	Options *config.MainConfig
}

func MakeShortanerHandler(c *config.MainConfig, db *storage.Database) ShortanerHandler {
	return ShortanerHandler{
		Db:      db,
		Options: c,
	}
}
