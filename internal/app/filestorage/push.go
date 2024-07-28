package filestorage

import (
	"encoding/json"

	"log/slog"
)

func (fs *Storage) Push(short string, original string) error {
	err := fs.cache.Push(short, original)
	if err != nil {
		slog.Error("Push to cache storage error", "err", err)
		return err
	}

	item := makeStorageItem(short, original)

	data, err := json.Marshal(item)
	if err != nil {
		slog.Error("When convert to json error", "err", err)
		return err
	}

	data = append(data, '\n')

	_, err = fs.file.Write(data)
	if err != nil {
		slog.Error("When write to file error", "err", err)
		return err
	}

	return nil
}
