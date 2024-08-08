package filestorage

import (
	"bufio"
	"encoding/json"
	"mishin-shortener/internal/app/mapstorage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("get_data_from_cache", func(t *testing.T) {
		db := mapstorage.Make()
		db.Push("short", "original")

		fs := Storage{cache: *db, file: nil}

		value, err := fs.Get("short")
		assert.Equal(t, value, "original")

		require.NoError(t, err)
	})
}

// данный тест проверяет, что происходит корректная запись в файл и причем так, что потом
// этот результат можно попарсить

func TestPush(t *testing.T) {
	t.Run("push_data_to_file", func(t *testing.T) {
		testFile, _ := os.CreateTemp("", "pattern")
		defer os.Remove(testFile.Name())

		fs := Make(testFile.Name()) // создаем fs
		fs.Push("short", "original")

		reader := bufio.NewReader(testFile)
		data, _ := reader.ReadBytes('\n')

		item := storageItem{}
		json.Unmarshal(data, &item)

		assert.Equal(t, item.OriginalURL, "original")
		assert.Equal(t, item.ShortURL, "short")
	})
}
