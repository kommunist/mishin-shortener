package middleware

import (
	"context"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

func AuthCheck(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		var authCookieValue string
		ctx := r.Context()

		authCookie, _ := r.Cookie("Authorization") // обработать ошибку
		if authCookie != nil {
			authCookieValue = authCookie.Value
		}

		if authCookieValue != "" { // если хедер с авторизацией есть
			userID, err := secure.Decrypt(authCookieValue)
			if err != nil || userID == "" { // и если не удалось расшифровать
				w.WriteHeader(http.StatusUnauthorized)
			} else { // единственный положительный сценарий
				ctx = context.WithValue(ctx, secure.UserIDKey, userID)
			}
		} else { // если хедера с авторизацией нет
			w.WriteHeader(http.StatusUnauthorized)
		}

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}
