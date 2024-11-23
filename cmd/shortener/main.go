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

	// go func() { // сервер для профилирования
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()

	h := app.Make()
	err := h.Call()
	if err != nil {
		slog.Error("Error in head package", "err", err)
		panic(err)
	}

}
