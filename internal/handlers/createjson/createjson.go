package createjson

import (
	"encoding/json"
	"io"
	"log/slog"
	"mishin-shortener/internal/errors/exist"
	"mishin-shortener/internal/hasher"
	"mishin-shortener/internal/secure"
	"net/http"
)

// Структура входящего запроса на сокращение в формате JSON.
type RequestItem struct {
	URL string `json:"url"`
}

// Структура ответа на сокращение в формате JSON.
type ResponseItem struct {
	Result string `json:"result"`
}

// Обработчик запроса на сокращение в формате JSON.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("Error when read body", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := RequestItem{}
	output := ResponseItem{}

	if err = json.Unmarshal(body, &input); err != nil {
		slog.Error("Parsing Error", "err", err)
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	hashed := hasher.GetMD5Hash([]byte(input.URL))
	err = h.storage.Push(r.Context(), hashed, string(input.URL), userID)

	status := http.StatusCreated

	if err != nil {
		if _, ok := err.(*exist.ExistError); ok { // обрабатываем проблему, когда уже есть в базе
			status = http.StatusConflict
		} else {
			slog.Error("push to storage error", "err", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	output.Result = string(h.resultURL(hashed))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	out, err := json.Marshal(output)
	if err != nil {
		slog.Error("error when encoding data", "err", err)
		http.Error(w, "Encoding json Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		slog.Error("error when write response", "err", err)
		http.Error(w, "Write response error", http.StatusInternalServerError)
		return
	}
}
