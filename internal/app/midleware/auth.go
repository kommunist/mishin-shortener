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
			slog.Info("UserID decrypted", "UserID", userID)

			if err != nil || userID == "" { // и если не удалось расшифровать
				slog.Warn("Error when decrypt header", "Header", authHeader)
				w.WriteHeader(http.StatusUnauthorized)
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
