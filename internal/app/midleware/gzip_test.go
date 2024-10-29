package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGzipHandler(w http.ResponseWriter, r *http.Request) {
}

func TestGzip(t *testing.T) {

	nextHandler := http.HandlerFunc(testGzipHandler)
	handlerToTest := Gzip(nextHandler)

	request :=
		httptest.NewRequest(http.MethodGet, "/any", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, request)

	res := w.Result()

	assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200 with auth")
}
