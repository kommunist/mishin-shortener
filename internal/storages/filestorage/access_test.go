package filestorage

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"os"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	exList := []struct {
		name        string
		fileName    string
		fileContent []storageItem
		short       string
		original    string
	}{
		{
			name:     "simple_test_happy_path",
			fileName: "tempFile.json",
			fileContent: []storageItem{
				makeStorageItem("short", "original"),
			},
			short:    "short",
			original: "original",
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			// создаем файлик с контентом
			if len(ex.fileContent) > 0 {
				file, _ := os.OpenFile(ex.fileName, os.O_RDWR|os.O_CREATE, 0666)
				for _, item := range ex.fileContent {
					content, _ := json.Marshal(item)
					file.Write(content)
					file.Write([]byte("\n"))
				}
				file.Close()
			}
			defer os.Remove(ex.fileName)

			stor, err := Make(ex.fileName)
			assert.NoError(t, err)
			defer stor.Finish()

			value, err := stor.Get(context.Background(), ex.short)
			assert.Equal(t, ex.original, value)
			assert.NoError(t, err)
		})
	}
}

func TestPush(t *testing.T) {
	exList := []struct {
		name        string
		fileName    string
		fileContent []storageItem
		short       string
		original    string
		pushError   bool
	}{
		{
			name:     "simple_test_happy_path",
			fileName: "tempFile.json",
			fileContent: []storageItem{
				makeStorageItem("short", "original"),
			},
			short:     "new_short",
			original:  "new_original",
			pushError: false,
		},
		{
			name:     "when_push_to_cache_return_error",
			fileName: "tempFile.json",
			fileContent: []storageItem{
				makeStorageItem("short", "original"),
			},
			short:     "new_short",
			original:  "new_original",
			pushError: true,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			// создаем файлик с контентом
			if len(ex.fileContent) > 0 {
				file, _ := os.OpenFile(ex.fileName, os.O_RDWR|os.O_CREATE, 0666)
				for _, item := range ex.fileContent {
					content, _ := json.Marshal(item)
					file.Write(content)
					file.Write([]byte("\n"))
				}
				file.Close()
			}
			defer os.Remove(ex.fileName)

			stor, err := Make(ex.fileName)
			assert.NoError(t, err)

			defer stor.Finish()

			if ex.pushError {
				ctrl := gomock.NewController(t)
				cacher := NewMockCacher(ctrl)
				stor.cache = cacher

				ctx := context.Background()

				returnedErr := errors.New("Some error")

				cacher.EXPECT().Push(ctx, ex.short, ex.original, "userID").Return(returnedErr)
				value := stor.Push(context.Background(), ex.short, ex.original, "userID")
				assert.Equal(t, returnedErr, value)
			} else {
				stor.Push(context.Background(), ex.short, ex.original, "userID")

				// Делаем простой тест, что оно легло в оперативный кеш
				value, err := stor.Get(context.Background(), ex.short)
				assert.Equal(t, ex.original, value)
				assert.NoError(t, err)
				stor.Finish()

				// создали базу на том же файле
				newStor, err := Make(ex.fileName)
				assert.NoError(t, err)
				value, err = newStor.Get(context.Background(), ex.short)
				assert.Equal(t, ex.original, value)
				assert.NoError(t, err)
			}
		})
	}
}

// TODO reafactor this
func TestPushBatch(t *testing.T) {
	t.Run("push_batch_data_to_file", func(t *testing.T) {
		testFile, _ := os.CreateTemp("", "pattern")
		defer os.Remove(testFile.Name())

		testData := make(map[string]string)
		testData["vupsen"] = "pupsen"
		testData["biba"] = "boba"

		fs, err := Make(testFile.Name()) // создаем fs
		assert.NoError(t, err)
		defer fs.Finish()
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
