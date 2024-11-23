package createjson

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mishin-shortener/internal/config"
	"mishin-shortener/internal/errors/exist"
	"mishin-shortener/internal/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateURLByJSON(t *testing.T) {
	exList := []struct {
		name          string
		inputJSON     []byte
		ctx           context.Context
		userID        string
		status        int
		original      string
		respBody      string
		hashed        string
		storErr       error
		storTimes     int
		checkRespBody bool
	}{
		{
			name:          "Start_POST_to_create_record_in_storage",
			inputJSON:     func() []byte { data, _ := json.Marshal(RequestItem{URL: "biba"}); return data }(),
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			status:        http.StatusCreated,
			original:      "biba",
			hashed:        "931691969b142b3a0f11a03e36fcc3b7",
			respBody:      "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
			storErr:       nil,
			storTimes:     1,
			checkRespBody: true,
		},
		{
			name:          "Start_POST_to_create_record_in_storage_when_already_exist",
			inputJSON:     func() []byte { data, _ := json.Marshal(RequestItem{URL: "biba"}); return data }(),
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			status:        http.StatusConflict,
			original:      "biba",
			hashed:        "931691969b142b3a0f11a03e36fcc3b7",
			respBody:      "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
			storErr:       exist.NewExistError(nil),
			storTimes:     1,
			checkRespBody: false,
		},
		{
			name:          "Start_POST_to_create_record_in_storage_when_without_user_id_in_context",
			inputJSON:     func() []byte { data, _ := json.Marshal(RequestItem{URL: "biba"}); return data }(),
			ctx:           context.Background(),
			userID:        "qq",
			status:        http.StatusInternalServerError,
			original:      "biba",
			hashed:        "931691969b142b3a0f11a03e36fcc3b7",
			respBody:      "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
			storErr:       nil,
			storTimes:     0,
			checkRespBody: false,
		},
		{
			name:          "Start_POST_to_create_record_in_storage_when_incorrect_input_json",
			inputJSON:     []byte("abracadabra"),
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			status:        http.StatusInternalServerError,
			original:      "biba",
			hashed:        "931691969b142b3a0f11a03e36fcc3b7",
			respBody:      "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
			storErr:       nil,
			storTimes:     0,
			checkRespBody: false,
		},
		{
			name:          "Start_POST_to_create_record_in_storage_when_another_error_from_stor",
			inputJSON:     func() []byte { data, _ := json.Marshal(RequestItem{URL: "biba"}); return data }(),
			ctx:           context.WithValue(context.Background(), secure.UserIDKey, "qq"),
			userID:        "qq",
			status:        http.StatusInternalServerError,
			original:      "biba",
			hashed:        "931691969b142b3a0f11a03e36fcc3b7",
			respBody:      "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
			storErr:       errors.New("qq"),
			storTimes:     1,
			checkRespBody: false,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockPusher(ctrl)

			c := config.MakeConfig()
			h := Make(c, stor)

			request := httptest.NewRequest(
				http.MethodPost, "/api/shorten", bytes.NewReader(ex.inputJSON),
			).WithContext(ex.ctx)

			stor.EXPECT().Push(
				ex.ctx, ex.hashed, ex.original, ex.userID,
			).Times(ex.storTimes).Return(ex.storErr)

			// Создаем рекорер, вызываем хендлер и сразу снимаем результат
			w := httptest.NewRecorder()
			h.Call(w, request)
			res := w.Result()

			// проверим статус ответа
			assert.Equal(t, ex.status, res.StatusCode)

			if ex.checkRespBody {
				// проверим содержимое ответа
				outputData := ResponseItem{}
				resBody, _ := io.ReadAll(res.Body)
				json.Unmarshal(resBody, &outputData)
				defer res.Body.Close()

				assert.Equal(t, ex.respBody, outputData.Result)
			}

		})

	}
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		c := config.MakeConfig()

		ctrl := gomock.NewController(b)
		defer ctrl.Finish()

		stor := NewMockPusher(ctrl)

		h := Make(c, stor)

		ctx := context.WithValue(context.Background(), secure.UserIDKey, "qq")

		stor.EXPECT().Push(ctx, "931691969b142b3a0f11a03e36fcc3b7", "biba", "qq")

		inputData := RequestItem{URL: "biba"}
		inputJSON, _ := json.Marshal(inputData)

		request :=
			httptest.NewRequest(
				http.MethodPost, "/api/shorten", bytes.NewReader(inputJSON),
			).WithContext(ctx)

		// Создаем рекорер, вызываем хендлер и сразу снимаем результат
		w := httptest.NewRecorder()
		b.StartTimer()
		h.Call(w, request)
	}
}
