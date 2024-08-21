package ch_1_server

import (
	"net/http"
	"testing"
)

func TestServers(t *testing.T) {
	server := http.Server{
		Addr: "localhost:5000", // memasukan address
	}

	// Menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
