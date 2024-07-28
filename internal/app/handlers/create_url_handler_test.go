package handlers

import (
	"io"
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/mapstorage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateURLHandler(t *testing.T) {
	t.Run("Start_POST_to_create_record_in_db", func(t *testing.T) {
		db := mapstorage.Make()
		c := config.MakeConfig()
		h := MakeShortanerHandler(c, &db)

		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru"))
		w := httptest.NewRecorder()
		h.CreateURLHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		require.NoError(t, err)

		assert.Equal(t, "http://localhost:8080/06509a58eff5d07b614ea9057d6c2a79", string(resBody))
	})
}
