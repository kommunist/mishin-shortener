package filestorage

import (
	"log/slog"
	"mishin-shortener/internal/app/mapstorage"
	"os"

	"github.com/google/uuid"
)

type Storage struct {
	cache mapstorage.Storage
	file  *os.File
}

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

	return &Storage{cache: cache, file: file}
}

type storageItem struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func makeStorageItem(short string, original string) storageItem {
	return storageItem{
		ShortURL:    short,
		OriginalURL: original,
		UUID:        uuid.New().String(),
	}
}
