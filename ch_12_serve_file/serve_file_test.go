package ch_12_serve_file

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		// ! serve file mengirimkan file sebagai response ke web
		http.ServeFile(writer, request, "./resources/named-html.html")
	} else {
		http.ServeFile(writer, request, "./resources/no-name.html")
	}
}

//go:embed resources/named-html.html
var named string

//go:embed resources/no-name.html
var noName string

func ServeFileEmbedHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprintln(writer, named)
	} else {
		fmt.Fprintln(writer, noName)
	}
}
func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(ServeFileHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(ServeFileEmbedHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
