package ch_6_query_parameter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name") // mengambil dat dari query parameter
	age := request.URL.Query().Get("age")
	if name == "" || age == "" {
		fmt.Fprintln(writer, "hello")
	} else {
		fmt.Fprintf(writer, "Hello %s im %s years old", name, age)
	}
}

func MultipleValuesHandler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query() // mengambil query
	names := query["name"]       // mendapatkan value untuk query name meskipun ada beberapa query name yang sama
	fmt.Fprint(writer, strings.Join(names, ", "))
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http:localhost:5000?name=Nurdin&age=20", nil)
	recorder := httptest.NewRecorder()

	Handler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestMultipleValueQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http:localhost:5000?name=Nurdin&name=Hishasy&name=Naruto&name=Sasuke", nil) // ada name yang duplikat
	recorder := httptest.NewRecorder()

	MultipleValuesHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
