package handlers

import (
	"mishin-shortener/internal/app/config"
	"mishin-shortener/internal/app/mapstorage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	t.Run("when_not_pg_storage_then_error", func(t *testing.T) {
		db := mapstorage.Make() // сделаем, что база не postgres
		c := config.MakeConfig()
		h := MakeShortanerHandler(c, db)

		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		w := httptest.NewRecorder()

		h.PingHandler(w, request)

		res := w.Result()
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

		defer res.Body.Close()
	})
}
