package filestorage

import (
	"context"
	"encoding/json"
	"log/slog"
)

func (fs *Storage) Get(ctx context.Context, shortURL string) (string, error) {

	return fs.cache.Get(ctx, shortURL)
}

func (fs *Storage) Push(ctx context.Context, short string, original string) error {
	err := fs.cache.Push(ctx, short, original)
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

func (fs *Storage) PushBatch(ctx context.Context, list *map[string]string) error {
	for k, v := range *list {
		err := fs.Push(ctx, k, v)
		if err != nil {
			slog.Error("When batch push to file error", "err", err)
			return err
		}

	}
	return nil
}

func (fs *Storage) Ping(ctx context.Context) error {
	return nil
}

func (fs *Storage) Finish() error {
	err := fs.file.Close()

	if err != nil {
		slog.Error("Failed close write to file", "err", err)
	}

	return err
}
