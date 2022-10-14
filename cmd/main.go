package main

import (
	"net/http"

	"github.com/wrik0/url-shortner/handlers"
)

var (
	shortnerHandler *handlers.Shortner
)

func init() {
	shortnerHandler = handlers.NewShortnerHandler()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/r/([^[a-zA-Z][0-9]{1,8}])", shortnerHandler.Handle)
	mux.HandleFunc("/ganja", shortnerHandler.Handle)
	// TODO remove
	http.ListenAndServe(":9000", mux)
}
