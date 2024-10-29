package middleware

import (
	"mishin-shortener/internal/app/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHandler(w http.ResponseWriter, r *http.Request) {}

func TestAuth(t *testing.T) {
	t.Run("correct_auth", func(t *testing.T) {

		userID := "user"
		encrypted, _ := secure.Encrypt(userID)

		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthCheck(nextHandler)

		request :=
			httptest.NewRequest(http.MethodGet, "/any", nil)

		request.AddCookie(
			&http.Cookie{
				Name:  "Authorization",
				Value: encrypted,
				Path:  "/",
			},
		)

		w := httptest.NewRecorder()
		handlerToTest.ServeHTTP(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200 with auth")
	})

	t.Run("incorrect_auth", func(t *testing.T) {

		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthCheck(nextHandler)

		request :=
			httptest.NewRequest(http.MethodGet, "/any", nil)

		request.AddCookie(
			&http.Cookie{
				Name:  "Authorization",
				Value: "abrkjsdfjknkjfnwkejnfwjenfjkwenkfacadabra",
				Path:  "/",
			},
		)

		w := httptest.NewRecorder()
		handlerToTest.ServeHTTP(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "response status must be 200 with auth")
	})

	t.Run("incorrect_auth_without_cookie", func(t *testing.T) {

		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthCheck(nextHandler)

		request :=
			httptest.NewRequest(http.MethodGet, "/any", nil)

		w := httptest.NewRecorder()
		handlerToTest.ServeHTTP(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "response status must be 200 with auth")
	})
}
