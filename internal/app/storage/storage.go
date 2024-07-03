package storage

type Database map[string]string

func (db *Database) Push(shortUrl string, originalUrl string) {
	(*db)[shortUrl] = originalUrl
}

func (db *Database) Get(shortUrl string) string {
	return (*db)[shortUrl]
}
