package mapstorage

import (
	"errors"
)

type Storage map[string]string

func Make() Storage {
	return Storage{}
}

func (db *Storage) Push(short string, original string) error {
	(*db)[short] = original

	return nil
}

func (db *Storage) Get(short string) (string, error) {
	value, ok := (*db)[short]
	if ok {
		return value, nil
	}

	return "", errors.New("not found")
}
