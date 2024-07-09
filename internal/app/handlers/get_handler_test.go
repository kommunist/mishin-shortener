package handlers

import (
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "testing"
    "mishin-shortener/internal/app/storage"
)

func TestGetHandler(t *testing.T) {

    t.Run("Start GET on not persisted url", func(t *testing.T) {
        database := storage.Database{} // пустая база

        request := httptest.NewRequest(http.MethodGet, "/qwerty", nil)
        w := httptest.NewRecorder()
        GetHandler(w, request, &database)

        res := w.Result()
        assert.Equal(t, http.StatusBadRequest, res.StatusCode)
    })

    t.Run("Start GET on url persisted in db", func(t *testing.T) {
        shorted := "/qwerty"
        expected := "ya.ru"

        database := storage.Database{shorted: expected} // пустая база

        request := httptest.NewRequest(http.MethodGet, shorted, nil)
        w := httptest.NewRecorder()
        GetHandler(w, request, &database)

        res := w.Result()
        assert.Equal(t, http.StatusTemporaryRedirect, res.StatusCode)
        assert.Equal(t, expected, res.Header.Get("Location"))
    })
}
