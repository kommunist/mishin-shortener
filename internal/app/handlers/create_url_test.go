package handlers

import (
	"context"
	"errors"
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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		stor.EXPECT().Push(
			ctx,
			"931691969b142b3a0f11a03e36fcc3b7",
			"biba",
			"qq",
		).Return(nil)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()

		h.CreateURL(w, request)

		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", string(resBody))
	})

	t.Run("Start_POST_to_create_record_in_db_when_exist", func(t *testing.T) {
		// создаём контроллер мока
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		stor.EXPECT().Push(
			ctx,
			"931691969b142b3a0f11a03e36fcc3b7",
			"biba",
			"qq",
		).Return(exsist.NewExistError(nil))

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()
		h.CreateURL(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusConflict, res.StatusCode)

		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7", string(resBody))
	})

	t.Run("Start_POST_to_create_record_in_db_when_another_error", func(t *testing.T) {
		// создаём контроллер мока
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		stor.EXPECT().Push(
			ctx,
			"931691969b142b3a0f11a03e36fcc3b7",
			"biba",
			"qq",
		).Return(errors.New("another"))

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()
		h.CreateURL(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("Start_POST_to_create_record_in_db_when_without_user_in_context", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		stor := mocks.NewMockAbstractStorage(ctrl)

		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, stor)

		ctx := context.Background()

		stor.EXPECT().Push(ctx, "931691969b142b3a0f11a03e36fcc3b7", "biba", "qq").Times(0)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()
		h.CreateURL(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}

func BenchmarkCreateURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		db := mapstorage.Make()
		c := config.MakeConfig()
		c.InitConfig()
		h := MakeShortanerHandler(c, db)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru")).WithContext(ctx)
		w := httptest.NewRecorder()
		b.StartTimer()
		h.CreateURL(w, request)
	}
}
