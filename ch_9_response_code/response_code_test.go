package ch_9_response_code

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		// Mengirimkan responsecode dengan writeheader
		writer.WriteHeader(400) // bad request
		fmt.Fprint(writer, "Name not provide")
	} else {
		writer.WriteHeader(200) // success
		fmt.Fprint(writer, "Hello ", name)
	}
}

func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	ResponseHandler(rec, req)

	res := rec.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))   //
	fmt.Println(res.StatusCode) // harusnya 400

	req = httptest.NewRequest(http.MethodGet, "http://localhost:5000?name=Nurdin", nil)
	rec = httptest.NewRecorder()

	ResponseHandler(rec, req)

	res = rec.Result()
	body, _ = io.ReadAll(res.Body)

	fmt.Println(string(body))   //
	fmt.Println(res.StatusCode) // harusnya 200
}
