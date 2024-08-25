package ch_18_template_caching

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ! Caching template menggunakan embed dan mengeluarkan template menjadi sebuah variable

//go:embed templates/*.html
var templates embed.FS

var myTemplate = template.Must(template.ParseFS(templates, "templates/*.html"))

func Caching(writer http.ResponseWriter, req *http.Request) {
	myTemplate.ExecuteTemplate(writer, "caching.html", map[string]any{
		"Name": "Nurdin",
	})
}

func TestCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	Caching(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
