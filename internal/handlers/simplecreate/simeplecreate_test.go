package simplecreate

import (
	"context"
	"errors"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/secure"
	"mishin-shortener/internal/errors/exsist"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name        string
		withContext bool
		storTimes   int
		storErr     error
		status      int
		respBody    string
	}{
		{
			name:        "post_simple_create_happy_path",
			withContext: true,
			storTimes:   1,
			storErr:     nil,
			status:      http.StatusCreated,
			respBody:    "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
		},
		{
			name:        "post_simple_create_when_without_user_in_context",
			withContext: false,
			storTimes:   0,
			storErr:     nil,
			status:      http.StatusInternalServerError,
			respBody:    "",
		},
		{
			name:        "post_simple_create_when_record_exist_in_db",
			withContext: true,
			storTimes:   1,
			storErr:     exsist.NewExistError(nil),
			status:      http.StatusConflict,
			respBody:    "http://localhost:8080/931691969b142b3a0f11a03e36fcc3b7",
		},
		{
			name:        "post_simple_create_when_another_error",
			withContext: true,
			storTimes:   1,
			storErr:     errors.New("another"),
			status:      http.StatusInternalServerError,
			respBody:    "",
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

			ctx := context.Background()
			if ex.withContext {
				ctx = context.WithValue(ctx, secure.UserIDKey, "qq")
			}

			stor.EXPECT().Push(
				ctx,
				"931691969b142b3a0f11a03e36fcc3b7",
				"biba",
				"qq",
			).Times(ex.storTimes).Return(ex.storErr)

			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
			w := httptest.NewRecorder()

			h.Call(w, request)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, ex.status, res.StatusCode)
			if len(ex.respBody) > 0 {
				respBody, _ := io.ReadAll(res.Body)
				assert.Equal(t, ex.respBody, string(respBody))
			}
		})
	}
}

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ctrl := gomock.NewController(b)
		defer ctrl.Finish()

		stor := NewMockPusher(ctrl)

		c := config.MakeConfig()
		c.InitConfig()
		h := Make(c, stor)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "qq")

		stor.EXPECT().Push(
			ctx,
			"931691969b142b3a0f11a03e36fcc3b7",
			"biba",
			"qq",
		).Times(1).Return(nil)

		request :=
			httptest.NewRequest(http.MethodPost, "/", strings.NewReader("biba")).WithContext(ctx)
		w := httptest.NewRecorder()
		b.StartTimer()
		h.Call(w, request)
	}
}
