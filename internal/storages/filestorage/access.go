package filestorage

import (
	"context"
	"encoding/json"
	"log/slog"
	"mishin-shortener/internal/delasync"
)

// Получение полного URL по сокращенному
func (fs *Storage) Get(ctx context.Context, shortURL string) (string, error) {

	return fs.cache.Get(ctx, shortURL)
}

// Сохранение в базу новой пары сокращенный/полный
func (fs *Storage) Push(ctx context.Context, short string, original string, userID string) error {
	err := fs.cache.Push(ctx, short, original, userID)
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

// Сохранение в базу списка пар сокращенный/полный
func (fs *Storage) PushBatch(ctx context.Context, list *map[string]string, userID string) error {
	for k, v := range *list {
		err := fs.Push(ctx, k, v, userID)
		if err != nil {
			slog.Error("When batch push to file error", "err", err)
			return err
		}

	}
	return nil
}

// Получение из базы списка сокращенных ссылок для пользователя(не реализовано)
func (fs *Storage) GetByUserID(ctx context.Context, userID string) (map[string]string, error) {
	return nil, nil
}

// Удаление из базы базы сокращенного URL для пользователя(не реализовано)
func (fs *Storage) DeleteByUserID(ctx context.Context, list []delasync.DelPair) error {
	return nil
}

// Восстановление коннектов к базе(не реализовано)
func (fs *Storage) Ping(ctx context.Context) error {
	return nil
}

// Завершение работы с хранилищем
func (fs *Storage) Finish() error {
	err := fs.file.Close()

	if err != nil {
		slog.Error("Failed close write to file", "err", err)
	}

	return err
}
