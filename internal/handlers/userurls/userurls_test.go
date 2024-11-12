package userurls

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/secure"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name       string
		withUserID bool
		storTimes  int
		storResult map[string]string
		storError  error
		status     int
		respItems  []responseItem
	}{
		{
			name:       "post_to_create_record_in_db",
			withUserID: true,
			storTimes:  1,
			storResult: map[string]string{"short0": "long0", "short1": "long1"},
			storError:  nil,
			status:     http.StatusOK,
			respItems: []responseItem{
				{
					Short:    "http://localhost:8080/short0",
					Original: "long0",
				},
				{
					Short:    "http://localhost:8080/short1",
					Original: "long1",
				},
			},
		},
		{
			name:       "post_to_create_record_in_db_without_userID_in_context",
			withUserID: false,
			storTimes:  0,
			storResult: nil,
			storError:  nil,
			status:     http.StatusInternalServerError,
			respItems:  nil,
		},
		{
			name:       "post_to_create_record_in_db_when_error_from_database",
			withUserID: true,
			storTimes:  1,
			storResult: nil,
			storError:  errors.New("some error"),
			status:     http.StatusInternalServerError,
			respItems:  nil,
		},
		{
			name:       "post_to_create_record_in_db_when_empty_result_from_db",
			withUserID: true,
			storTimes:  1,
			storResult: map[string]string{},
			storError:  nil,
			status:     http.StatusNoContent,
			respItems:  nil,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			stor := NewMockByUserIDGetter(ctrl)

			ctx := context.Background()
			if ex.withUserID {
				ctx = context.WithValue(ctx, secure.UserIDKey, "userId")
			}

			stor.EXPECT().GetByUserID(
				ctx,
				"userId",
			).Times(ex.storTimes).Return(ex.storResult, ex.storError)

			c := config.MakeConfig()
			c.InitConfig()
			h := Make(c, stor)

			req := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil).WithContext(ctx)

			w := httptest.NewRecorder()
			h.Call(w, req)
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, ex.status, res.StatusCode)

			if len(ex.respItems) > 0 {
				respList := make([]responseItem, 0)
				resBody, _ := io.ReadAll(res.Body)
				json.Unmarshal(resBody, &respList)

				for _, item := range ex.respItems {
					assert.Contains(t, respList, item, "response contain correct result")
				}
			}
		})
	}
}

func BenchmarkUserURLs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		ctrl := gomock.NewController(b)
		stor := NewMockByUserIDGetter(ctrl)

		ctx := context.Background()
		ctx = context.WithValue(ctx, secure.UserIDKey, "userId")

		// используется не для проверки, а для мака базы
		stor.EXPECT().GetByUserID(
			ctx,
			"userId",
		).Return(map[string]string{"short0": "long0", "short1": "long1"}, nil)

		c := config.MakeConfig()
		c.InitConfig()
		h := Make(c, stor)

		request := httptest.NewRequest(http.MethodGet, "/api/user/urls", nil).WithContext(ctx)

		w := httptest.NewRecorder()
		b.StartTimer()
		h.Call(w, request)
	}
}
