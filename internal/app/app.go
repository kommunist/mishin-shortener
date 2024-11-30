package app

import (
	"log"
	"log/slog"
	"mishin-shortener/internal/api/grpcserver"
	"mishin-shortener/internal/api/httpserver"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/delasync"
	"mishin-shortener/internal/storages/filestorage"
	"mishin-shortener/internal/storages/mapstorage"
	"mishin-shortener/internal/storages/pgstorage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type finisher interface {
	Finish() error
}

type item struct {
	storage finisher
	setting config.MainConfig
	HTTPAPI *httpserver.HTTPHandler
	GRPCAPI *grpcserver.GRPCHandler

	deleter delasync.Handler
}

type commonStorage interface {
	httpserver.CommonStorage // методы для httpApi
	delasync.Remover         // методы для worker
	finisher                 // методы для закрытия
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
func Make() (item, error) {
	c := config.MakeConfig()
	err := c.InitConfig()
	if err != nil {
		slog.Error("Error from InitConfig")
		return item{}, err
	}

	storage := initStorage(c)     // создали хранилище
	del := delasync.Make(storage) // создали асинхронный обработчик удалений

	h := httpserver.Make(c, storage, del.DelChan)
	g := grpcserver.Make(c, storage, del.DelChan)

	return item{HTTPAPI: h, GRPCAPI: g, deleter: del, storage: storage, setting: c}, nil
}

// Основной метод объекта item
func (i *item) Call() error {

	if i.setting.EnableProfile {
		i.startProfileServer()
	}

	i.listenInterrupt()
	i.deleter.InitWorker()

	go func() {
		i.GRPCAPI.Start()
	}()

	err := i.HTTPAPI.Call()
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
	i.HTTPAPI.Stop()
	i.GRPCAPI.Stop()

	err := i.storage.Finish()
	if err != nil {
		slog.Error("Error when winish storage", "err", err)
	}

	close(sigint)
}

func (i *item) startProfileServer() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
