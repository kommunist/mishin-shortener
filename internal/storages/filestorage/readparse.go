package filestorage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func readAndParse(file *os.File) ([]storageItem, error) {
	reader := bufio.NewReader(file)
	list := []storageItem{}

	for {
		data, eof := reader.ReadBytes('\n')

		if eof == io.EOF {
			slog.Info("Full file readed.", "num_items_founded", len(list))
			break
		}

		item := storageItem{}

		err := json.Unmarshal(data, &item)
		fmt.Println(string(data))
		if err != nil {
			fmt.Println(err)
			slog.Error("input file JSON parsing error", "err", err)
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}
