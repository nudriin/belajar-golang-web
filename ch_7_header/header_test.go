package ch_7_header

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HeaderHandler(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writter, request.Header.Get("content-type"))  // mendapatkan content type
	fmt.Fprintln(writter, request.Header.Get("authorization")) // mendapatkan content type
	writter.Header().Add("user", "Nurdin")                     // membuat response header
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	request.Header.Add("Content-Type", "application/json")                                   // Set header di request
	request.Header.Set("Authorization", "Bearer dasd8d1nsd8nw19nw19nd1891nwdsmmx8m8f9m2908") // set header

	recorder := httptest.NewRecorder()

	HeaderHandler(recorder, request)

	// Membaca response body
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body) // Byte

	fmt.Println(string(body))
	fmt.Println(recorder.Header().Get("user")) // mengambil response header

}
