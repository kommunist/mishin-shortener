package handlers

import (
	"encoding/json"
	"io"
	"mishin-shortener/internal/app/hasher"
	"net/http"
)

type RequestBatchItem struct {
	CorrelationId string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ResponseBatchItem struct {
	CorrelationId string `json:"correlation_id"`
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

		prepareToSave[v.OriginalURL] = "/" + hashed

		output = append(
			output,
			ResponseBatchItem{CorrelationId: v.CorrelationId, ShortURL: hashed},
		)
	}

	err = h.DB.PushBatch(&prepareToSave)
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
