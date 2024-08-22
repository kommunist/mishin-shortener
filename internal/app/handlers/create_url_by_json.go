package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"mishin-shortener/internal/app/exsist"
	"mishin-shortener/internal/app/hasher"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

type RequestData struct {
	URL string `json:"url"`
}

type ResponseData struct {
	Result string `json:"result"`
}

func (h *ShortanerHandler) CreateURLByJSON(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		slog.Error("Error when read body", "err", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := RequestData{}
	output := ResponseData{}

	if err = json.Unmarshal(body, &input); err != nil {
		slog.Error("Parsing Error", "err", err)
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		userID = ""
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	hashed := hasher.GetMD5Hash([]byte(input.URL))
	err = h.DB.Push(r.Context(), "/"+hashed, string(input.URL), userID)

	status := http.StatusCreated

	if err != nil {
		if _, ok := err.(*exsist.ExistError); ok { // обрабатываем проблему, когда уже есть в базе
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

	w.Write(out)
}
