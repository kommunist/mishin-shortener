package filestorage

func (fs *Storage) Get(shortURL string) (string, error) {

	return fs.cache.Get(shortURL)
}
