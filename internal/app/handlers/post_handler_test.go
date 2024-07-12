package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"mishin-shortener/internal/app/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostHandler(t *testing.T) {
	t.Run("Start POST to create record in db", func(t *testing.T) {
		database := storage.Database{} // пустая база

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru"))
		w := httptest.NewRecorder()
		PostHandler(w, request, &database, "http://example.com")

		res := w.Result()
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://example.com/06509a58eff5d07b614ea9057d6c2a79", string(resBody))
	})
}
