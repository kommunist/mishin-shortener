package filestorage

import (
	"bufio"
	"context"
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
		db.Push(context.Background(), "short", "original", "userID")

		fs := Storage{cache: *db, file: nil}

		value, err := fs.Get(context.Background(), "short")
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
		fs.Push(context.Background(), "short", "original", "userID")

		reader := bufio.NewReader(testFile)
		data, _ := reader.ReadBytes('\n')

		item := storageItem{}
		json.Unmarshal(data, &item)

		assert.Equal(t, item.OriginalURL, "original")
		assert.Equal(t, item.ShortURL, "short")
	})
}

// данный тест проверяет, что происходит корректная запись в файл и причем так, что потом
// этот результат можно попарсить

func TestPushBatch(t *testing.T) {
	t.Run("push_batch_data_to_file", func(t *testing.T) {
		testFile, _ := os.CreateTemp("", "pattern")
		defer os.Remove(testFile.Name())

		testData := make(map[string]string)
		testData["vupsen"] = "pupsen"
		testData["biba"] = "boba"

		fs := Make(testFile.Name()) // создаем fs
		fs.PushBatch(context.Background(), &testData, "userID")

		reader := bufio.NewReader(testFile)
		data, _ := reader.ReadBytes('\n')

		firstItem := storageItem{}
		json.Unmarshal(data, &firstItem)

		data, _ = reader.ReadBytes('\n')

		secondItem := storageItem{}
		json.Unmarshal(data, &secondItem)

		assert.Contains(t, []string{"pupsen", "boba"}, firstItem.OriginalURL)
		assert.Contains(t, []string{"pupsen", "boba"}, secondItem.OriginalURL)

		assert.Contains(t, []string{"vupsen", "biba"}, firstItem.ShortURL)
		assert.Contains(t, []string{"vupsen", "biba"}, secondItem.ShortURL)
	})
}
