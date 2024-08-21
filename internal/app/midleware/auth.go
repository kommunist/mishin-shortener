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
		var userID string

		authHeader := r.Header.Get("Authorization")

		if authHeader != "" { // если хедер с авторизацией есть
			var err error
			slog.Info("Authorization founded", "auth", authHeader)

			userID, err = secure.Decrypt(authHeader)
			if err != nil { // и если не удалось расшифровать
				slog.Warn("Error when decrypt header", "Header", authHeader)
				userID = newuserID() // создадим новый
			}
			if userID != "" {
				slog.Info("User id decrypted from header", "Id", userID)
			}
		} else { // если хедера с авторизацией нет
			userID = newuserID()
		}

		ctx := context.WithValue(r.Context(), "userID", userID)

		encryptedId, _ := secure.Encrypt(userID) // сделать обработку ошибки
		w.Header().Set("Authorization", encryptedId)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}

func newuserID() string {
	return uuid.New().String()
}
