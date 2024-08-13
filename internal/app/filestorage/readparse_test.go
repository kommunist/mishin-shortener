package filestorage

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// данный тест проверяет, что корректно парсятся список записей из файла

func TestReadAndParse(t *testing.T) {
	t.Run("read_and_parse_data_from_file", func(t *testing.T) {
		testFile, _ := os.CreateTemp("", "pattern")
		defer os.Remove(testFile.Name())

		fs := Make(testFile.Name()) // создаем fs
		fs.Push(context.Background(), "short0", "original0")
		fs.Push(context.Background(), "short1", "original1")

		list := readAndParse(testFile)

		assert.Equal(t, len(list), 2)

		assert.Equal(t, list[0].OriginalURL, "original0")
		assert.Equal(t, list[0].ShortURL, "short0")

		assert.Equal(t, list[1].OriginalURL, "original1")
		assert.Equal(t, list[1].ShortURL, "short1")
	})
}
