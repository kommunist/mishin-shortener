package mapstorage

import (
	"context"
	"errors"
	"log/slog"
)

type Storage map[string]string

func Make() *Storage {
	return &Storage{}
}

func (db *Storage) Push(ctx context.Context, short string, original string, userID string) error {
	(*db)[short] = original

	return nil
}

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

func (db *Storage) Get(ctx context.Context, short string) (string, error) {
	value, ok := (*db)[short]
	if ok {
		return value, nil
	}

	return "", errors.New("not found")
}

func (db *Storage) GetByUserID(ctx context.Context, userID string) (map[string]string, error) {
	return nil, nil
}

func (db *Storage) DeleteByUserID(ctx context.Context, userID string, shorts []string) error {
	return nil
}

func (db *Storage) Ping(ctx context.Context) error {
	return nil
}

func (db *Storage) Finish() error {
	return nil
}
