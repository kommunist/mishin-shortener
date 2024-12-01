package createjsonbatch

import (
	"encoding/json"
	"io"
	"log/slog"
	"mishin-shortener/internal/secure"
	"net/http"
)

// Обработчик запроса на сокращение в формате JSON батчами.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := []requestBatchItem{}

	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	output, err := h.Perform(r.Context(), input, userID)
	if err != nil {
		http.Error(w, "Error when perform service", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		slog.Error("error when write response", "err", err)
		http.Error(w, "Write response error", http.StatusInternalServerError)
		return
	}
}
