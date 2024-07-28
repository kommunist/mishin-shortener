package filestorage

import (
	"encoding/json"
	"log"
)

func (fs *Storage) Push(short string, original string) error {
	err := fs.cache.Push(short, original)
	if err != nil {
		log.Println(err)
		return err
	}

	item := makeStorageItem(short, original)

	data, err := json.Marshal(item)
	if err != nil {
		log.Println(err)
		return err
	}

	data = append(data, '\n')

	_, err = fs.file.Write(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
