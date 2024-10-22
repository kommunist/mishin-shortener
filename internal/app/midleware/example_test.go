package middleware

import "net/http"

func ExampleAuthCheck() {
	testHandler := func(w http.ResponseWriter, r *http.Request) {}

	nextHandler := http.HandlerFunc(testHandler)
	AuthCheck(nextHandler)
}

func ExampleAuthSet() {
	testHandler := func(w http.ResponseWriter, r *http.Request) {}

	nextHandler := http.HandlerFunc(testHandler)
	AuthCheck(nextHandler)
}
