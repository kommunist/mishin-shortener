package handlers

import (
	"context"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/exsist"
	"mishin-shortener/internal/app/mapstorage"
	"mishin-shortener/internal/app/secure"
	"mishin-shortener/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateURL(t *testing.T) {
	t.Run("Start_POST_to_create_record_in_db", func(t *testing.T) {
		db := mapstorage.Make()
		c := config.MakeConfig()
		h := MakeShortanerHandler(c, db)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIdKey, "qq")

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru")).WithContext(ctx)
		w := httptest.NewRecorder()
		h.CreateURL(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://localhost:8080/06509a58eff5d07b614ea9057d6c2a79", string(resBody))
	})

	t.Run("Start_POST_to_create_record_in_db_when_exist", func(t *testing.T) {
		// создаём контроллер мока
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIdKey, "qq")

		stor.EXPECT().Push(
			ctx,
			"/931691969b142b3a0f11a03e36fcc3b7",
			"biba",
			"qq",
		).Return(exsist.NewExistError(nil))

		c := config.MakeConfig()
		h := MakeShortanerHandler(c, stor)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()
		h.CreateURL(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusConflict, res.StatusCode)

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", string(resBody))
	})
}
