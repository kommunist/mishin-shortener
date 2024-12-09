package middleware

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGzipHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Response"))
}

func TestGzip(t *testing.T) {
	exList := []struct {
		name                string
		withCorrectCompress bool
	}{
		{
			name:                "when_all_succee",
			withCorrectCompress: true,
		},
		{
			name:                "when_with_error",
			withCorrectCompress: false,
		},
	}
	for _, ex := range exList {
		t.Run(ex.name, func(t *testing.T) {

			nextHandler := http.HandlerFunc(testGzipHandler)
			handlerToTest := Gzip(nextHandler)

			var buf bytes.Buffer

			if ex.withCorrectCompress {
				g := gzip.NewWriter(&buf)
				g.Write([]byte("qqq"))

			} else {
				buf.Write([]byte("qqq"))
			}

			request :=
				httptest.NewRequest(http.MethodGet, "/any", &buf)

			request.Header.Set("Accept-Encoding", "gzip")
			request.Header.Set("Content-Encoding", "gzip")

			w := httptest.NewRecorder()
			handlerToTest.ServeHTTP(w, request)

			res := w.Result()
			defer res.Body.Close()

			if ex.withCorrectCompress {
				assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200")
			} else {
				assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "response status must be 500")
			}

		})

	}

}
