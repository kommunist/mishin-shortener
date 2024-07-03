package main

import (
	"io"
	"net/http"

	"internal/hasher"
)

func createURLHandler(postResponse http.ResponseWriter, postRequest *http.Request) {
	if postRequest.Method != http.MethodPost {
		http.Error(postResponse, "Bad Request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(postRequest.Body)
	if err != nil {
		panic(err)
	}

	hashed := hasher.GetMD5Hash(body)

	http.HandleFunc(
		"/"+hashed, func(getResponse http.ResponseWriter, getRequest *http.Request) {
			if getRequest.Method != http.MethodGet {
				http.Error(getResponse, "Bad Request", http.StatusBadRequest)
			} else {
				getResponse.Header().Set("Location", string(body))
				getResponse.WriteHeader(http.StatusTemporaryRedirect)
			}
		})
	postResponse.WriteHeader(http.StatusCreated)
	postResponse.Write([]byte("http://" + postRequest.Host + "/" + hashed))
}

func main() {
	http.HandleFunc("/", createURLHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
