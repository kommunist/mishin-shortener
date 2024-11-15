// Пакет main - основной для запуска приложения
package main

import (
	"log"
	"log/slog"

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
	err := c.InitConfig()
	if err != nil {
		slog.Error("Error from InitConfig")
		panic(err)
	}

	storage := initStorage(c)

	defer func() {
		defErr := storage.Finish()
		if defErr != nil {
			slog.Error("Error when finish with storage", "err", defErr)
		}
	}()

	a := api.Make(c, storage)
	err = a.Call()
	if err != nil {
		slog.Error("Error from api component", "err", err)
		panic(err)
	}
}
