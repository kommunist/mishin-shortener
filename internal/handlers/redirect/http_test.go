package redirect

import (
	"context"
	"errors"
	"mishin-shortener/internal/errors/deleted"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	exList := []struct {
		name           string
		shorted        string
		expected       string
		expectedStatus int
		returnedError  error
	}{
		{
			name:           "Start_GET_on_url_persisted_in_db",
			shorted:        "qwerty",
			expected:       "ya.ru",
			expectedStatus: http.StatusTemporaryRedirect,
			returnedError:  nil,
		},
		{
			name:           "Start_GET_on_url_already_deleted",
			shorted:        "qqqq",
			expected:       "",
			expectedStatus: http.StatusGone,
			returnedError:  deleted.NewDeletedError(nil),
		},
		{
			name:           "Start_GET_on_another_error",
			shorted:        "qqqq",
			expected:       "",
			expectedStatus: http.StatusNotFound,
			returnedError:  errors.New("ququ"),
		},
	}

	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stor := NewMockGetter(ctrl)
			h := Make(stor)

			ctx := context.Background()

			stor.EXPECT().Get(ctx, ex.shorted).Return(ex.expected, ex.returnedError)

			request := httptest.NewRequest(
				http.MethodGet, "/"+ex.shorted, nil,
			).WithContext(ctx)

			w := httptest.NewRecorder()
			h.Call(w, request)

			res := w.Result()
			res.Body.Close()
			assert.Equal(t, ex.expectedStatus, res.StatusCode)
			assert.Equal(t, ex.expected, res.Header.Get("Location"))
		})

	}
}
