package ch_8_post_form

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostFormHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // parsing form yang dikirim dari client
	if err != nil {
		panic(err)
	}

	name := request.PostForm.Get("name")
	age := request.PostForm.Get("age")

	fmt.Fprintf(writer, "my name is %s and im %s years old", name, age)
}

func PostFormHandler2(writer http.ResponseWriter, request *http.Request) {

	// Bisa juga seperti ini, auto parsing dari function PostFormValue
	name := request.PostFormValue("name")
	age := request.PostFormValue("age")

	fmt.Fprintf(writer, "my name is %s and im %s years old", name, age)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("name=Nurdin&age=20") // membuat requst body
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5000", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") // mengubah content typenya

	recorder := httptest.NewRecorder()

	PostFormHandler(recorder, request)
	// Membaca response body
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body) // Byte

	fmt.Println(string(body))
}
