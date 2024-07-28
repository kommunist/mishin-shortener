package filestorage

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

func readAndParse(file *os.File) []storageItem {
	reader := bufio.NewReader(file)
	list := []storageItem{}

	for {
		data, eof := reader.ReadBytes('\n')

		item := storageItem{}

		err := json.Unmarshal(data, &item)
		if err != nil {
			log.Fatalf("input file JSON parsing error")
		}
		list = append(list, item)

		if eof == io.EOF {
			log.Printf("Full file readed. Founded %d items", len(list))
			break
		}
	}

	return list
}
