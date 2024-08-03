package handlers

import (
	"log/slog"
	"mishin-shortener/internal/app/config"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

type AbstractStorage interface {
	Push(string, string) error
	Get(string) (string, error)
	Finish() error
}

type ShortanerHandler struct {
	DB      AbstractStorage
	Options config.MainConfig
	Driver  *sql.DB
}

func MakeShortanerHandler(c config.MainConfig, db AbstractStorage) ShortanerHandler {
	driver, err := sql.Open("postgres", c.DatabaseDSN)

	if err != nil {
		slog.Error("Eror when connect to database", "err", err)
		os.Exit(1)
	}

	slog.Info("Success connect to database")

	return ShortanerHandler{
		DB:      db,
		Options: c,
		Driver:  driver,
	}
}
