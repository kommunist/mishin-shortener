// Пакет main - основной для запуска приложения
package main

import (
	"log"
	"log/slog"
	"os"

	"mishin-shortener/internal/api"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/filestorage"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/app/mapstorage"
	"mishin-shortener/internal/app/pgstorage"

	"net/http"

	_ "net/http/pprof"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"
var buildCommit string = "N/A"

func initStorage(c config.MainConfig) handlers.AbstractStorage {
	if c.DatabaseDSN != "" {
		return pgstorage.Make(c)
	}
	if c.FileStoragePath != "" {
		return filestorage.Make(c.FileStoragePath)
	}

	return mapstorage.Make()
}

func main() {
	slog.Info("Build info", "version", buildVersion)
	slog.Info("Build info", "date", buildDate)
	slog.Info("Build info", "commit", buildCommit)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	c := config.MakeConfig()
	c.InitConfig()

	storage := initStorage(c)

	defer func() {
		err := storage.Finish()
		if err != nil {
			slog.Error("Error when finish with storage", "err", err)
		}
	}()

	a := api.Make(c, storage)
	err := a.Call()
	if err != nil {
		slog.Error("Error from api component", "err", err)
		os.Exit(1)
	}
}
