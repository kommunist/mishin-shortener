package middleware

import "net/http"

func ExampleAuthCheck() {
	testHandlerAuth := func(w http.ResponseWriter, r *http.Request) {}

	nextHandler := http.HandlerFunc(testHandlerAuth)
	AuthCheck(nextHandler)
}

func ExampleAuthSet() {
	testHandlerAuth := func(w http.ResponseWriter, r *http.Request) {}

	nextHandler := http.HandlerFunc(testHandlerAuth)
	AuthCheck(nextHandler)
}
