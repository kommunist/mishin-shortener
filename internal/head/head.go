package head

import (
	"fmt"
	"log/slog"
	"mishin-shortener/internal/api"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/handlers"
	"mishin-shortener/internal/storages/filestorage"
	"mishin-shortener/internal/storages/mapstorage"
	"mishin-shortener/internal/storages/pgstorage"
	"os"
	"os/signal"
	"syscall"
)

type item struct {
	API *api.ShortanerAPI
}

func initStorage(c config.MainConfig) handlers.AbstractStorage {
	if c.DatabaseDSN != "" {
		return pgstorage.Make(c)
	}
	if c.FileStoragePath != "" {
		return filestorage.Make(c.FileStoragePath)
	}

	return mapstorage.Make()
}

// Конструктор объекта item
func Make() item {
	c := config.MakeConfig()

	err := c.InitConfig()
	if err != nil {
		slog.Error("Error from InitConfig")
		panic(err)
	}
	storage := initStorage(c)
	fmt.Println("qq")
	a := api.Make(c, storage)

	return item{API: &a}
}

// Основной метод объекта item
func (i *item) Call() {
	i.listenInterrupt()

	err := i.API.Call()
	if err != nil {
		slog.Error("Error from api component", "err", err)
		panic(err)
	}
}

func (i *item) listenInterrupt() { // регистрируем канал для прерываний и перенаправляем туда внешние прерывания
	sigint := make(chan os.Signal, 3)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go i.waitInterrupt(sigint)
}

func (i *item) waitInterrupt(sigint chan os.Signal) {
	<-sigint // ждем сигнал прeрывания

	i.API.Stop()

	close(sigint)
}
