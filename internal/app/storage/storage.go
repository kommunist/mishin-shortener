package storage

type Database map[string]string

func (db *Database) Push(shortURL string, originalURL string) {
	(*db)[shortURL] = originalURL
}

func (db *Database) Get(shortURL string) string {
	return (*db)[shortURL]
}
