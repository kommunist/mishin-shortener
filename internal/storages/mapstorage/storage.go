// Модуль mapstorage предоставляет хранение данных в MAP структуре.
package mapstorage

import (
	"context"
	"errors"
	"log/slog"
	"mishin-shortener/internal/delasync"
)

// Структура хранения данных.
type Storage map[string]string

// Создает структуру хранения данных.
func Make() *Storage {
	return &Storage{}
}

// Сохранение в базу новой пары сокращенный/полный
func (db *Storage) Push(ctx context.Context, short string, original string, userID string) error {
	(*db)[short] = original

	return nil
}

// Сохранение в базу списка пар сокращенный/полный
func (db *Storage) PushBatch(ctx context.Context, list *map[string]string, userID string) error {
	for k, v := range *list {
		err := db.Push(ctx, k, v, userID)
		if err != nil {
			slog.Error("When batch push to mapstorage error", "err", err)
			return err
		}

	}
	return nil
}

// Получение полного URL по сокращенному
func (db *Storage) Get(ctx context.Context, short string) (string, error) {
	value, ok := (*db)[short]
	if ok {
		return value, nil
	}

	return "", errors.New("not found")
}

// Получение из базы списка сокращенных ссылок для пользователя(не реализовано)
func (db *Storage) GetByUserID(ctx context.Context, userID string) (map[string]string, error) {
	return nil, nil
}

// Удаление из базы базы сокращенного URL для пользователя(не реализовано)
func (db *Storage) DeleteByUserID(ctx context.Context, list []delasync.DelPair) error {
	return nil
}

// Получение статистики о кол-ве пользователей и урлов в базе
func (db *Storage) GetStats(ctx context.Context) (int, int, error) {
	return 0, 0, nil
}

// Восстановление коннектов к базе(не реализовано)
func (db *Storage) Ping(ctx context.Context) error {
	return nil
}

// Завершение работы с хранилищем
func (db *Storage) Finish() error {
	return nil
}
