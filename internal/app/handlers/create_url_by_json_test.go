package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/exsist"
	"mishin-shortener/internal/app/mapstorage"
	"mishin-shortener/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateURLByJSON(t *testing.T) {
	t.Run("Start_POST_to_create_record_in_storage", func(t *testing.T) {
		db := mapstorage.Make()
		c := config.MakeConfig()
		h := MakeShortanerHandler(c, db)

		inputData := RequestData{URL: "biba"}
		inputJSON, _ := json.Marshal(inputData)

		request :=
			httptest.NewRequest(
				http.MethodPost,
				"/api/shorten",
				bytes.NewReader(inputJSON),
			)

		// Создаем рекорер, вызываем хендлер и сразу снимаем результат
		w := httptest.NewRecorder()
		h.CreateURLByJSON(w, request)
		res := w.Result()

		// проверим статус ответа
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		// проверим содержимое ответа
		outputData := ResponseData{}
		resBody, _ := io.ReadAll(res.Body)
		json.Unmarshal(resBody, &outputData)
		defer res.Body.Close()

		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", outputData.Result)

		// проверим содержимое базы
		var v string
		v, _ = db.Get("/931691969b142b3a0f11a03e36fcc3b7")
		assert.Equal(t, "biba", v)
	})

	t.Run("Start_POST_to_create_record_in_storage_when_already_exist", func(t *testing.T) {
		// создаём контроллер мока
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)
		stor.EXPECT().Push("/931691969b142b3a0f11a03e36fcc3b7", "biba").Return(exsist.NewExistError(nil))

		c := config.MakeConfig()
		h := MakeShortanerHandler(c, stor)

		inputData := RequestData{URL: "biba"}
		inputJSON, _ := json.Marshal(inputData)

		request :=
			httptest.NewRequest(
				http.MethodPost,
				"/api/shorten",
				bytes.NewReader(inputJSON),
			)

		// Создаем рекорер, вызываем хендлер и сразу снимаем результат
		w := httptest.NewRecorder()
		h.CreateURLByJSON(w, request)
		res := w.Result()

		// проверим статус ответа
		assert.Equal(t, http.StatusConflict, res.StatusCode)

		// проверим содержимое ответа
		outputData := ResponseData{}
		resBody, _ := io.ReadAll(res.Body)
		json.Unmarshal(resBody, &outputData)
		defer res.Body.Close()

		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", outputData.Result)
	})
}
