package handlers

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
    "internal/storage"
    "strings"
)

func TestStatusHandler(t *testing.T) {

    t.Run("Start GET on not persisted url", func(t *testing.T) {
        database := storage.Database{} // пустая база

        request := httptest.NewRequest(http.MethodGet, "/qwerty", nil)
        w := httptest.NewRecorder()
        MainHandler(w, request, &database)

        res := w.Result()
        assert.Equal(t, http.StatusBadRequest, res.StatusCode)
    })

    t.Run("Start GET on url persisted in db", func(t *testing.T) {
        shorted := "/qwerty"
        expected := "ya.ru"

        database := storage.Database{shorted: expected} // пустая база

        request := httptest.NewRequest(http.MethodGet, shorted, nil)
        w := httptest.NewRecorder()
        MainHandler(w, request, &database)

        res := w.Result()
        assert.Equal(t, http.StatusTemporaryRedirect, res.StatusCode)
        assert.Equal(t, expected, res.Header.Get("Location"))
    })

    t.Run("Start POST to create record in db", func(t *testing.T) {
      database := storage.Database{} // пустая база

      request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ya.ru"))
      w := httptest.NewRecorder()
      MainHandler(w, request, &database)

      res := w.Result()
      assert.Equal(t, http.StatusCreated, res.StatusCode)

      defer res.Body.Close()
      resBody, err := io.ReadAll(res.Body)
      require.NoError(t, err)

      assert.Equal(t, "http://example.com/06509a58eff5d07b614ea9057d6c2a79", string(resBody))
    })
}
