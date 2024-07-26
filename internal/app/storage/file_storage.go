package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"mishin-shortener/internal/app/config"
	"os"
)

type FileStorage struct {
	cache CacheStorage
	file  *os.File
}

func MakeFileStorage(c config.MainConfig) FileStorage {
	var file *os.File
	cache := MakeCacheStorage()

	if c.FileStoragePath != "" {
		openedFile, err := os.OpenFile(c.FileStoragePath, os.O_RDWR|os.O_CREATE, 0666)
		file = openedFile

		if err != nil {
			fmt.Println("FATAL")
		}
		items := readAndParse(file)
		for _, v := range items {
			cache[v.ShortURL] = v
		}
		fmt.Printf("Readed %d items", len(cache))
	}

	return FileStorage{cache: cache, file: file}
}

func (fs *FileStorage) Push(short string, original string) {
	fs.cache.Push(short, original)

	itemStruct, err := fs.cache.GetItem(short)
	if err != nil {
		fmt.Println("FATAL")
	}

	data, err := json.Marshal(itemStruct)
	if err != nil {
		fmt.Println("FATAL")
	}

	data = append(data, '\n')

	_, err = fs.file.Write(data)
	if err != nil {
		fmt.Println("Fatal")
	}
}

func (fs *FileStorage) Get(shortURL string) (string, error) {

	return fs.cache.Get(shortURL)
}

func readAndParse(file *os.File) []storageItem {
	reader := bufio.NewReader(file)
	list := []storageItem{}

	for {
		data, eof := reader.ReadBytes('\n')

		item := storageItem{}

		err := json.Unmarshal(data, &item)
		if err != nil {
			fmt.Println("FATAL2")
		}
		list = append(list, item)

		if eof == io.EOF {
			fmt.Println("READED FULL")
			break
		}
	}

	return list
}
