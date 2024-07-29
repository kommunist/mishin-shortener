package filestorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// данный тест проверяет, что происходит корректный пар запись в файл и причем так, что потом
// этот результат можно попарсить

func TestReadAndParse(t *testing.T) {
	t.Run("read_and_parse_data_from_file", func(t *testing.T) {
		testFile, _ := os.CreateTemp("", "pattern")
		defer os.Remove(testFile.Name())

		fs := Make(testFile.Name()) // создаем fs
		fs.Push("short0", "original0")
		fs.Push("short1", "original1")

		list := readAndParse(testFile)

		assert.Equal(t, len(list), 2)

		assert.Equal(t, list[0].OriginalURL, "original0")
		assert.Equal(t, list[0].ShortURL, "short0")

		assert.Equal(t, list[1].OriginalURL, "original1")
		assert.Equal(t, list[1].ShortURL, "short1")
	})
}
