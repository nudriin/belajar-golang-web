package ch_19_auto_escape

import (
	"html/template"
	"net/http"
	"testing"
)

func AutoEscapeDisabled(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "template.html", map[string]any{
		"Title": "Auto Escape OFF",
		"Body":  template.HTML("<p>Selamat Belajar Golang</p>"), // Mematikan auto escape
	})
}

func TestAutoEscape(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(AutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
