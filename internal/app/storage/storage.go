package storage

import "errors"

type Database map[string]string

func MakeDatabase() Database {
	return Database{}
}

func (db *Database) Push(shortURL string, originalURL string) {
	(*db)[shortURL] = originalURL
}

func (db *Database) Get(shortURL string) (string, error) {
	value := (*db)[shortURL]

	if value == "" {
		return "", errors.New("not found")
	}

	return value, nil
}
