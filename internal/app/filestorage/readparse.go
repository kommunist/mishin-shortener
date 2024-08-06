package filestorage

import (
	"bufio"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

func readAndParse(file *os.File) []storageItem {
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
		if err != nil {
			slog.Error("input file JSON parsing error", "err", err)
			os.Exit(1)
		}
		list = append(list, item)
	}

	return list
}
