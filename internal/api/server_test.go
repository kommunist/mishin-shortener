package api

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/storages/mapstorage"
	"testing"
	"time"
)

func TestCall(t *testing.T) {
	// данный тест проверяет непосредственно запуск сервера и
	// что вся предварительная работа с роутерами проходит без проблем
	// и сервер запускается
	t.Run("happy_path_router_inited", func(t *testing.T) {
		setting := config.MakeConfig()
		stor := mapstorage.Make() // используем самый простой mapstorage

		api := Make(setting, stor)
		go func() {
			time.Sleep(1 * time.Second)
			api.Stop()
		}()
		api.Call()
	})
}
