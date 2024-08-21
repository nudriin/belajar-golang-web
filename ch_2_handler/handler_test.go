package ch_2_handler

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	// membuat handler
	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World") // print ke writter
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
