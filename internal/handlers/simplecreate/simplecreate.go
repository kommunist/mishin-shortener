package simplecreate

import (
	"io"
	"log/slog"
	"mishin-shortener/internal/app/hasher"
	"mishin-shortener/internal/app/secure"
	"mishin-shortener/internal/errors/exsist"
	"net/http"
)

// Обработчик простого запроса на сокращение.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("read body error", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	status := http.StatusCreated
	hashed := hasher.GetMD5Hash(body)

	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	err = h.storage.Push(r.Context(), hashed, string(body), userID)
	if err != nil {
		if _, ok := err.(*exsist.ExistError); ok { // обрабатываем проблему, когда уже есть в базе
			status = http.StatusConflict
		} else {
			slog.Error("push to storage error", "err", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return

		}
	}

	w.WriteHeader(status)
	_, err = w.Write(h.resultURL(hashed))
	if err != nil {
		slog.Error("error when write response", "err", err)
		http.Error(w, "Write response error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) resultURL(hashed string) []byte {
	return []byte(h.setting.BaseRedirectURL + "/" + hashed)
}
