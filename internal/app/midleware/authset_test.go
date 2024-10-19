package middleware

import (
	"mishin-shortener/internal/app/secure"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAuthHandler(w http.ResponseWriter, r *http.Request) {}

func TestAuthSet(t *testing.T) {
	t.Run("correct_auth", func(t *testing.T) {

		userId := "user"
		encrypted, _ := secure.Encrypt(userId)

		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthSet(nextHandler)

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

		decrypted, _ := secure.Decrypt(res.Cookies()[0].Value)

		assert.Equal(t, userId, decrypted, "response has same cookie as request")
		assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200 with auth")
	})

	t.Run("when_incorrect_cookie_value", func(t *testing.T) {
		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthSet(nextHandler)

		request :=
			httptest.NewRequest(http.MethodGet, "/any", nil)

		request.AddCookie(
			&http.Cookie{
				Name:  "Authorization",
				Value: "abracadwljfb;wbejfqabra",
				Path:  "/",
			},
		)

		w := httptest.NewRecorder()
		handlerToTest.ServeHTTP(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, "Authorization", res.Cookies()[0].Name, "response has cookie with Authorization name")
		assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200 with auth")
	})

	t.Run("set_auth_cookie_with_correct_name", func(t *testing.T) {

		nextHandler := http.HandlerFunc(testHandler)
		handlerToTest := AuthSet(nextHandler)

		request :=
			httptest.NewRequest(http.MethodGet, "/any", nil)

		w := httptest.NewRecorder()
		handlerToTest.ServeHTTP(w, request)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, "Authorization", res.Cookies()[0].Name, "response has cookie with Authorization name")
		assert.Equal(t, http.StatusOK, res.StatusCode, "response status must be 200")
	})

}
