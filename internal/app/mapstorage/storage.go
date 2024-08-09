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

func (db *Storage) Push(short string, original string) error {
	(*db)[short] = original

	return nil
}

func (db *Storage) PushBatch(list *map[string]string) error {
	for k, v := range *list {
		err := db.Push(k, v)
		if err != nil {
			slog.Error("When batch push to mapstorage error", "err", err)
			return err
		}

	}
	return nil
}

func (db *Storage) Get(short string) (string, error) {
	value, ok := (*db)[short]
	if ok {
		return value, nil
	}

	return "", errors.New("not found")
}

func (db *Storage) Ping(ctx context.Context) error {
	return nil
}

func (db *Storage) Finish() error {
	return nil
}
