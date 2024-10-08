package handlers

import (
	"context"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/mapstorage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name           string
		shorted        string
		expected       string
		expectedStatus int
		handler        ShortanerHandler
		beforeFunction func(*ShortanerHandler, string, string)
	}{
		{
			name:           "Start_GET_on_url_persisted_in_db",
			shorted:        "qwerty",
			expected:       "ya.ru",
			expectedStatus: http.StatusTemporaryRedirect,
			handler: func() ShortanerHandler {
				c := config.MakeConfig()
				db := mapstorage.Make()
				return MakeShortanerHandler(c, db)
			}(),
			beforeFunction: func(h *ShortanerHandler, shorted string, expected string) {
				h.DB.Push(context.Background(), shorted, expected, "userID")
			},
		},
		{
			name:           "Start_GET_on_url_not_persisted_in_db",
			shorted:        "qqqq",
			expected:       "",
			expectedStatus: http.StatusNotFound,
			handler: func() ShortanerHandler {
				c := config.MakeConfig()
				db := mapstorage.Make()
				return MakeShortanerHandler(c, db)
			}(),
		},
	}

	for _, testItem := range tests {
		t.Run(testItem.name, func(t *testing.T) {
			if testItem.beforeFunction != nil {
				testItem.beforeFunction(&testItem.handler, testItem.shorted, testItem.expected)
			}

			request := httptest.NewRequest(http.MethodGet, "/"+testItem.shorted, nil)
			w := httptest.NewRecorder()
			testItem.handler.RedirectHandler(w, request)

			res := w.Result()
			res.Body.Close()
			assert.Equal(t, testItem.expectedStatus, res.StatusCode)
			assert.Equal(t, testItem.expected, res.Header.Get("Location"))
		})

	}
}
