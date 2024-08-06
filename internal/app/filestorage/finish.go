package filestorage

import (
	"log/slog"
)

func (fs *Storage) Finish() error {
	err := fs.file.Close()

	if err != nil {
		slog.Error("Failed close write to file", "err", err)
	}

	return err
}
