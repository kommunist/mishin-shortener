package filestorage

import "github.com/google/uuid"

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
