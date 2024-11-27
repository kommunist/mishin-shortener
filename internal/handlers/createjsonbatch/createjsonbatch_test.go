package createjsonbatch

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateURLByJSONBatch(t *testing.T) {
	exList := []struct {
		name          string
		inputJSON     []byte
		ctx           context.Context
		userID        string
		status        int
		checkBody     bool
		responseItems []ResponseBatchItem
		storTimes     int
	}{
		{
			name: "Start_POST_to_create_record_in_storage",
			inputJSON: func() []byte {
				data, _ := json.Marshal([]RequestBatchItem{
					{CorrelationID: "123", OriginalURL: "biba"},
					{CorrelationID: "456", OriginalURL: "boba"},
				})
				return data
			}(),
			ctx:       context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:    "qq",
			status:    http.StatusCreated,
			checkBody: true,
			responseItems: []ResponseBatchItem{
				{CorrelationID: "123", ShortURL: "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7"},
				{CorrelationID: "456", ShortURL: "http://localhost:8080/2cce0ec300cfe8dd3024939db0448893"},
			},
			storTimes: 1,
		},
		{
			name: "Start_POST_to_create_record_in_storage_without_user_in_context",
			inputJSON: func() []byte {
				data, _ := json.Marshal([]RequestBatchItem{
					{CorrelationID: "123", OriginalURL: "biba"},
					{CorrelationID: "456", OriginalURL: "boba"},
				})
				return data
			}(),
			ctx:           context.Background(),
			userID:        "qq",
			status:        http.StatusInternalServerError,
			checkBody:     false,
			responseItems: nil,
			storTimes:     0,
		},
		{
			name:          "Start_POST_to_create_record_in_storage_when_input_data_broken",
			inputJSON:     []byte("abracadabra"),
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			status:        http.StatusInternalServerError,
			checkBody:     false,
			responseItems: nil,
			storTimes:     0,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockPusher(ctrl)

			c := config.MakeConfig()
			c.InitConfig()

			h := Make(c, stor)

			stor.EXPECT().PushBatch(
				ex.ctx,
				&map[string]string{
					"931691969b142b3a0f11a03e36fcc3b7": "biba", "2cce0ec300cfe8dd3024939db0448893": "boba",
				},
				ex.userID,
			).Times(ex.storTimes)

			request :=
				httptest.NewRequest(
					http.MethodPost, "/api/shorten/batch", bytes.NewReader(ex.inputJSON),
				).WithContext(ex.ctx)

			// Создаем рекорер, вызываем хендлер и сразу снимаем результат
			w := httptest.NewRecorder()
			h.Call(w, request)
			res := w.Result()
			defer res.Body.Close()

			// проверим статус ответа
			assert.Equal(t, ex.status, res.StatusCode)

			if ex.checkBody {
				// проверим содержимое ответа
				outputData := make([]ResponseBatchItem, 0, 2)

				resBody, _ := io.ReadAll(res.Body)
				json.Unmarshal(resBody, &outputData)
				defer res.Body.Close()

				for i, item := range ex.responseItems {
					assert.Equal(t, item.CorrelationID, outputData[i].CorrelationID)
					assert.Equal(t, item.ShortURL, outputData[i].ShortURL)
				}

			}

		})
	}
}

func BenchmarkCreateURLByJSONBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		c := config.MakeConfig()
		c.InitConfig()

		ctrl := gomock.NewController(b)
		defer ctrl.Finish()

		stor := NewMockPusher(ctrl)
		h := Make(c, stor)

		inputData := []RequestBatchItem{
			{CorrelationID: "123", OriginalURL: "biba"},
			{CorrelationID: "456", OriginalURL: "boba"},
		}
		inputJSON, _ := json.Marshal(inputData)

		ctx := context.WithValue(context.Background(), secure.UserIDKey, "qq")

		stor.EXPECT().PushBatch(
			ctx,
			&map[string]string{
				"931691969b142b3a0f11a03e36fcc3b7": "biba", "2cce0ec300cfe8dd3024939db0448893": "boba",
			},
			"qq",
		)

		request :=
			httptest.NewRequest(
				http.MethodPost,
				"/api/shorten/batch",
				bytes.NewReader(inputJSON),
			).WithContext(ctx)

		// Создаем рекорер, вызываем хендлер и сразу снимаем результат
		w := httptest.NewRecorder()
		b.StartTimer()
		h.Call(w, request)
	}

}
