package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testLogHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("teapot"))
}

func TestWithLogRequestResponse(t *testing.T) {

	nextHandler := http.HandlerFunc(testLogHandler)
	handlerToTest := WithLogRequest(nextHandler)

	request :=
		httptest.NewRequest(http.MethodGet, "/any", nil)

	w := httptest.NewRecorder()
	handlerToTest.ServeHTTP(w, request)

	res := w.Result()

	assert.Equal(t, http.StatusTeapot, res.StatusCode, "response status must be 200 with auth")
}
