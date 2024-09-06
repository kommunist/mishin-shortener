package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/mapstorage"
	"mishin-shortener/internal/app/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateURLByJSONBatch(t *testing.T) {
	t.Run("Start_POST_to_create_record_in_storage", func(t *testing.T) {
		db := mapstorage.Make()
		c := config.MakeConfig()
		h := MakeShortanerHandler(c, db)

		inputData := []RequestBatchItem{
			{CorrelationID: "123", OriginalURL: "biba"},
			{CorrelationID: "456", OriginalURL: "boba"},
		}
		inputJSON, _ := json.Marshal(inputData)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		request :=
			httptest.NewRequest(
				http.MethodPost,
				"/api/shorten/batch",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		// Создаем рекорер, вызываем хендлер и сразу снимаем результат
		w := httptest.NewRecorder()
		h.CreateURLByJSONBatch(w, request)
		res := w.Result()

		// проверим статус ответа
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		// проверим содержимое ответа
		outputData := make([]ResponseBatchItem, 0, 2)
		resBody, _ := io.ReadAll(res.Body)
		json.Unmarshal(resBody, &outputData)
		defer res.Body.Close()

		assert.Equal(t, "123", outputData[0].CorrelationID)
		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", outputData[0].ShortURL)
		assert.Equal(t, "456", outputData[1].CorrelationID)
		assert.Equal(t, "http://localhost:8080/2cce0ec300cfe8dd3024939db0448893", outputData[1].ShortURL)

		// проверим содержимое базы
		var v string
		v, _ = db.Get(context.Background(), "931691969b142b3a0f11a03e36fcc3b7")
		assert.Equal(t, "biba", v)

		v, _ = db.Get(context.Background(), "2cce0ec300cfe8dd3024939db0448893")
		assert.Equal(t, "boba", v)
	})
}
