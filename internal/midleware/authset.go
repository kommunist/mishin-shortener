package middleware

import (
	"context"
	"mishin-shortener/internal/secure"
	"net/http"

	"github.com/google/uuid"
)

// Мидлварь, который "создает" аутентификацию для noname пользователя
func AuthSet(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		var userID string
		var authCookieValue string

		authCookie, _ := r.Cookie("Authorization")
		if authCookie != nil {
			authCookieValue = authCookie.Value
		}

		if authCookieValue != "" { // если хедер с авторизацией есть
			var err error

			userID, err = secure.Decrypt(authCookieValue)
			if err != nil || userID == "" { // и если не удалось расшифровать
				userID = newuserID()
			}
		} else { // если хедера с авторизацией нет
			userID = newuserID()
		}

		ctx := context.WithValue(r.Context(), secure.UserIDKey, userID)

		encryptedID, _ := secure.Encrypt(userID)

		newCookie := newAuthCookie(encryptedID)
		http.SetCookie(w, &newCookie)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}

func newAuthCookie(value string) http.Cookie {
	return http.Cookie{
		Name:  "Authorization",
		Value: value,
		Path:  "/",
	}
}

func newuserID() string {
	return uuid.New().String()
}
