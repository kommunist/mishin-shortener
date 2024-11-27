package ping

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	exList := []struct {
		name     string
		returned error
		status   int
	}{
		{
			name:     "ping_storage_when_all_success",
			returned: nil,
			status:   http.StatusOK,
		},
		{
			name:     "ping_storage_when_ping_return_error",
			returned: errors.New("ququ"),
			status:   http.StatusInternalServerError,
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()

			stor := NewMockPinger(ctrl)
			h := Make(stor)

			stor.EXPECT().Ping(ctx).Return(ex.returned)

			request := httptest.NewRequest(http.MethodGet, "/ping", nil).WithContext(ctx)
			w := httptest.NewRecorder()

			h.Call(w, request)
			res := w.Result()
			assert.Equal(t, ex.status, res.StatusCode)

			defer res.Body.Close()
		})

	}
}
