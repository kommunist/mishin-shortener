package filestorage

import (
	"fmt"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/mapstorage"
	"os"

	"github.com/google/uuid"
)

type Storage struct {
	cache mapstorage.Storage
	file  *os.File
}

func MakeStorage(c config.MainConfig) Storage {
	var file *os.File
	cache := mapstorage.Make()

	if c.FileStoragePath != "" {
		openedFile, err := os.OpenFile(c.FileStoragePath, os.O_RDWR|os.O_CREATE, 0666)
		file = openedFile

		if err != nil {
			fmt.Println("FATAL")
		}
		items := readAndParse(file)
		for _, v := range items {
			cache[v.ShortURL] = v.OriginalURL
		}
		fmt.Printf("Readed %d items", len(cache))
	}

	return Storage{cache: cache, file: file}
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
