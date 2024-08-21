package ch_3_serve_mux

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	// ServeMux adalah handle yang mendukung multiple endpoint
	// membuat mux
	mux := http.NewServeMux()
	// Membuat handle func untuk endpoint tertentu
	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World") // print ke writter
	})

	mux.HandleFunc("/hello", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello you") // print ke writter
	})

	mux.HandleFunc("/hai", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "hai too") // print ke writter
	})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux, // memasukan handlernya
	}

	err := server.ListenAndServe() // memulai server

	if err != nil {
		panic(err)
	}
}
