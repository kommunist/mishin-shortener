package app

import (
	"log/slog"
	"mishin-shortener/internal/api"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/storages/filestorage"
	"mishin-shortener/internal/storages/mapstorage"
	"mishin-shortener/internal/storages/pgstorage"
	"os"
	"os/signal"
	"syscall"
)

type finisher interface {
	Finish() error
}

type item struct {
	storage finisher
	API     *api.ShortanerAPI
	deleter delasync.Handler
}

type commonStorage interface {
	api.CommonStorage // методы для api
	delasync.Remover  // методы для worker
	finisher          // методы для закрытия
}

func initStorage(c config.MainConfig) commonStorage {
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

	storage := initStorage(c)     // создали хранилище
	del := delasync.Make(storage) // создали асинхронный обработчик удалений

	a := api.Make(c, storage, del.DelChan)

	return item{API: a, deleter: del, storage: storage}
}

// Основной метод объекта item
func (i *item) Call() error {
	i.listenInterrupt()

	i.deleter.InitWorker()

	err := i.API.Call()
	if err != nil {
		slog.Error("Error from api component", "err", err)
		return err
	}
	return nil
}

func (i *item) listenInterrupt() { // регистрируем канал для прерываний и перенаправляем туда внешние прерывания
	sigint := make(chan os.Signal, 3)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	go i.waitInterrupt(sigint)
}

func (i *item) waitInterrupt(sigint chan os.Signal) {
	<-sigint // ждем сигнал прeрывания

	i.deleter.Stop()
	i.API.Stop()

	err := i.storage.Finish()
	if err != nil {
		slog.Error("Error when winish storage", "err", err)
	}

	close(sigint)
}
