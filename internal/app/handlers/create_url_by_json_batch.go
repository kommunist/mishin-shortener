package handlers

import (
	"encoding/json"
	"io"
	"mishin-shortener/internal/app/hasher"
	"mishin-shortener/internal/app/secure"
	"net/http"
)

type RequestBatchItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ResponseBatchItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

func (h *ShortanerHandler) CreateURLByJSONBatch(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := []RequestBatchItem{}
	output := []ResponseBatchItem{}

	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	prepareToSave := make(map[string]string)

	for _, v := range input {
		hashed := hasher.GetMD5Hash([]byte(v.OriginalURL))

		prepareToSave[hashed] = v.OriginalURL

		output = append(
			output,
			ResponseBatchItem{
				CorrelationID: v.CorrelationID,
				ShortURL:      h.Options.BaseRedirectURL + "/" + hashed,
			},
		)
	}

	var userID string
	if r.Context().Value(secure.UserIDKey) == nil {
		userID = ""
	} else {
		userID = r.Context().Value(secure.UserIDKey).(string)
	}

	err = h.DB.PushBatch(r.Context(), &prepareToSave, userID)

	if err != nil {
		http.Error(w, "Error when push to storage", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	w.Write(out)
}
