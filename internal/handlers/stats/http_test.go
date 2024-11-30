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

func TestCallHTTP(t *testing.T) {
	exList := []struct {
		name      string
		storTimes int
		storErr   error
		storUsers int
		storUrls  int
		status    int
		checkBody bool
		realIP    string
		subnet    string
	}{
		{
			name:      "simple_happy_path",
			storTimes: 1,
			storErr:   nil,
			storUsers: 1,
			storUrls:  1,
			status:    http.StatusOK,
			checkBody: true,
			realIP:    "192.168.1.1",
			subnet:    "192.168.1.0/24",
		},
		{
			name:      "when_error_from_db",
			storTimes: 1,
			storErr:   errors.New("Ququ"),
			storUsers: 0,
			storUrls:  0,
			status:    http.StatusInternalServerError,
			checkBody: false,
			realIP:    "192.168.1.1",
			subnet:    "192.168.1.0/24",
		},
		{
			name:      "when_ip_not_in_subnet",
			storTimes: 0,
			storErr:   errors.New("Ququ"),
			storUsers: 0,
			storUrls:  0,
			status:    http.StatusForbidden,
			checkBody: false,
			realIP:    "192.168.2.1",
			subnet:    "192.168.1.0/24",
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			stor := NewMockStatsGetter(ctrl)

			c := config.MakeConfig()
			c.InitConfig()
			c.TrustedSubnet = ex.subnet

			h, err := Make(c, stor)
			assert.NoError(t, err)

			ctx := context.Background()

			stor.EXPECT().GetStats(ctx).Times(ex.storTimes).Return(ex.storUsers, ex.storUrls, ex.storErr)

			request := httptest.NewRequest(http.MethodGet, "/", nil).WithContext(ctx)
			request.Header.Set("X-Real-IP", ex.realIP)

			w := httptest.NewRecorder()

			h.CallHTPP(w, request)

			res := w.Result()
			defer res.Body.Close()
			assert.Equal(t, ex.status, res.StatusCode)

			if ex.checkBody {
				respBody, _ := io.ReadAll(res.Body)

				item := responseItem{}
				json.Unmarshal(respBody, &item)
				assert.Equal(t, ex.storUrls, item.Urls)
				assert.Equal(t, ex.storUsers, item.Users)

			}
		})
	}

}
