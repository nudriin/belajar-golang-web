package ch16_template_layout

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func LayoutHandler(writer http.ResponseWriter, req *http.Request) {
	// ! Jika menggunakan parse file harus seperti ini (ribet)
	// t, err := template.ParseFiles(
	// 	"./templates/header.html",
	// 	"./templates/footer.html",
	// 	"./templates/content-layout.html"
	// )

	// ! Jika menggunakan parse glob lebih simple
	t, err := template.ParseGlob("./templates/*.html") // biar langsung di parse semua
	if err != nil {
		panic(err)
	}

	// ! Menggunakan "layout" yang namanya sudah di define
	t.ExecuteTemplate(writer, "layout", map[string]any{
		"Name": "Nurdin",
	})
}

func TestLayout(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	LayoutHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
