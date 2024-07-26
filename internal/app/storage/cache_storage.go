package storage

import (
	"errors"

	"github.com/google/uuid"
)

type storageItem struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type CacheStorage map[string]storageItem

func MakeCacheStorage() CacheStorage {
	return CacheStorage{}
}

func (db *CacheStorage) Push(short string, original string) {
	(*db)[short] = storageItem{UUID: uuid.New().String(), ShortURL: short, OriginalURL: original}
}

func (db *CacheStorage) Get(short string) (string, error) {
	value, ok := (*db)[short]
	if ok {
		return value.OriginalURL, nil
	}

	return "", errors.New("not found")
}

func (db *CacheStorage) GetItem(short string) (storageItem, error) {
	value, ok := (*db)[short]
	if ok {
		return value, nil
	}

	return storageItem{}, errors.New("not found")
}
