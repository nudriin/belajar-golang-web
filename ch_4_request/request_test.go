package ch_4_request

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	// membuat handler
	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World ")
		fmt.Fprint(writter, request.Method)     // Mengambil request method
		fmt.Fprint(writter, request.RequestURI) // Mengambil data request
	}

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: handler, // memasukan handlernya
	}

	err := server.ListenAndServe() // memulai server

	if err != nil {
		panic(err)
	}
}
