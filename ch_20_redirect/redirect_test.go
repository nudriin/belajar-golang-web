package ch_20_redirect

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectToHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello you redirect here")
}

func RedirectFromHandler(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/to", http.StatusTemporaryRedirect) // meridirect
}

func TestRedirect(t *testing.T) {

	mux := http.NewServeMux()

	mux.HandleFunc("/from", RedirectFromHandler)
	mux.HandleFunc("/to", RedirectToHandler)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
