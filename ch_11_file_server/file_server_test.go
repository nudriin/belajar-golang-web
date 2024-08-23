package ch_11_file_server

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")     // membuat path untuk directori
	fileServer := http.FileServer(dir) // membuat file server dengan param directories

	mux := http.NewServeMux()
	//! StripPrefix digunakan untuk menghapus path "/static" pada file jadi akan dibaca /resources/index.html
	//! jika tanpa StripPrefix maka akan dibaca /resources/static/index.html dan akan not found
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	dir, _ := fs.Sub(resources, "resources")    // masuk ke dalam folder resources yang ada pada embed
	fileServer := http.FileServer(http.FS(dir)) // menginputkan embed file menggunakan directories

	mux := http.NewServeMux()
	//! StripPrefix digunakan untuk menghapus path "/static" pada file jadi akan dibaca /resources/index.html
	//! jika tanpa StripPrefix maka akan dibaca /resources/static/index.html dan akan not found
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
