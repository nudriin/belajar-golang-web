package ch_5_http_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writter, "Hello world")
}

func TestHelloHandler(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:5000", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	// Membaca response body
	response := recorder.Result()
	body, err := io.ReadAll(response.Body) // Byte
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
