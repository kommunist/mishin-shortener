package stats

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mishin-shortener/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name      string
		storTimes int
		storErr   error
		storUsers int
		storUrls  int
		status    int
		checkBody bool
	}{
		{
			name:      "simple_happy_path",
			storTimes: 1,
			storErr:   nil,
			storUsers: 1,
			storUrls:  1,
			status:    http.StatusOK,
			checkBody: true,
		},
		{
			name:      "when_error_from_db",
			storTimes: 1,
			storErr:   errors.New("Ququ"),
			storUsers: 0,
			storUrls:  0,
			status:    http.StatusInternalServerError,
			checkBody: false,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			stor := NewMockStatsGetter(ctrl)

			c := config.MakeConfig()
			c.InitConfig()

			h, err := Make(c, stor)
			assert.NoError(t, err)

			ctx := context.Background()

			stor.EXPECT().GetStats(ctx).Times(ex.storTimes).Return(ex.storUsers, ex.storUrls, ex.storErr)

			request := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
			w := httptest.NewRecorder()

			h.Call(w, request)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, ex.status, res.StatusCode)

			if ex.checkBody {
				respBody, _ := io.ReadAll(res.Body)

				item := ResponseItem{}
				json.Unmarshal(respBody, &item)
				assert.Equal(t, ex.storUrls, item.Urls)
				assert.Equal(t, ex.storUsers, item.Users)

			}
		})
	}

}