// Пакет main - основной для запуска приложения
package main

import (
	"log/slog"
	"mishin-shortener/internal/app"

	_ "net/http/pprof"
)

var buildVersion string = "N/A"
var buildDate string = "N/A"
var buildCommit string = "N/A"

func main() {
	slog.Info("Build info", "version", buildVersion)
	slog.Info("Build info", "date", buildDate)
	slog.Info("Build info", "commit", buildCommit)

	h, err := app.Make()
	if err != nil {
		slog.Error("Error in make app package", "err", err)
		panic(err)
	}
	err = h.Call()
	if err != nil {
		slog.Error("Error in call app package", "err", err)
		panic(err)
	}
}
