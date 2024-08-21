package middleware

import (
	"context"
	"log/slog"
	"mishin-shortener/internal/app/secure"
	"net/http"

	"github.com/google/uuid"
)

func AuthMiddleware(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		var userId string

		authHeader := r.Header.Get("Authorization")

		if authHeader != "" { // если хедер с авторизацией есть
			var err error
			slog.Info("Authorization founded", "auth", authHeader)

			userId, err = secure.Decrypt(authHeader)
			if err != nil { // и если не удалось расшифровать
				slog.Warn("Error when decrypt header", "Header", authHeader)
				userId = newUserId() // создадим новый
			}
			if userId != "" {
				slog.Info("User id decrypted from header", "Id", userId)
			}
		} else { // если хедера с авторизацией нет
			userId = newUserId()
		}

		ctx := context.WithValue(r.Context(), "UserId", userId)

		encryptedId, _ := secure.Encrypt(userId) // сделать обработку ошибки
		w.Header().Set("Authorization", encryptedId)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}

func newUserId() string {
	return uuid.New().String()
}
