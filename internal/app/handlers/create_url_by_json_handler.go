package handlers

import (
	"encoding/json"
	"io"
	"mishin-shortener/internal/app/hasher"
	"net/http"
)

type RequestData struct {
	URL string `json:"url"`
}

type ResponseData struct {
	Result string `json:"result"`
}

func (h *ShortanerHandler) CreateURLByJSONHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	input := RequestData{}
	output := ResponseData{}

	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	hashed := hasher.GetMD5Hash([]byte(input.URL))

	h.DB.Push("/"+hashed, string(input.URL))

	output.Result = h.Options.BaseRedirectURL + "/" + hashed

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, "Parsing Error", http.StatusInternalServerError)
		return
	}

	w.Write(out)
}
