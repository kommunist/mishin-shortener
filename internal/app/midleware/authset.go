package middleware

import (
	"context"
	"mishin-shortener/internal/app/secure"
	"net/http"

	"github.com/google/uuid"
)

func AuthSet(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		var userID string
		ctx := r.Context()

		authHeader := r.Header.Get("Authorization")

		if authHeader != "" { // если хедер с авторизацией есть
			var err error

			userID, err = secure.Decrypt(authHeader)
			if err != nil || userID == "" { // и если не удалось расшифровать
				userID = newuserID()
			} else { // если расшифровался, то сохраним
				ctx = context.WithValue(r.Context(), "userID", userID)
			}
		} else { // если хедера с авторизацией нет
			userID = newuserID()
		}

		encryptedId, _ := secure.Encrypt(userID) // сделать обработку ошибки
		w.Header().Set("Authorization", encryptedId)

		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(authFn)
}

func newuserID() string {
	return uuid.New().String()
}
