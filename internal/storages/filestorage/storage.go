package filestorage

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/storages/mapstorage"
	"os"
)

// Интерфейс кешера для оперативной работы с хранилищем
type Cacher interface {
	Get(context.Context, string) (string, error)        // context, short
	Push(context.Context, string, string, string) error // context, short, original, userID
}

// Основная структура хранилища
type Storage struct {
	cache Cacher
	file  *os.File
}

// Функция создания хранилища
func Make(filePath string) *Storage {
	cache := *mapstorage.Make()

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		slog.Error("open file error", "err", err)
		os.Exit(1)
	}
	items, err := readAndParse(file)
	if err != nil {
		slog.Error("error when parse file", "err", err)
		os.Exit(1)
	}
	for _, v := range items {
		cache[v.ShortURL] = v.OriginalURL
	}
	slog.Info("readed n items", "n", len(cache))

	return &Storage{cache: &cache, file: file}
}
