package ch_14_template_data

import (
	"html/template"
	"net/http"
	"testing"
)

type DataTemplate struct {
	Title string
	Name  string
}

func FileTemplate(writer http.ResponseWriter, request *http.Request) {

	// Membuat data template menggunakan struct
	data := DataTemplate{
		Title: "Template Data",
		Name:  "Nurdin",
	}

	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "template.html", data) // memasukan data temmplate
}

func FileTemplate2(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	// ! Bisa juga menggunakan map seperti ini
	t.ExecuteTemplate(writer, "template.html", map[string]interface{}{
		"Title": "Template Data",
		"Name":  "Nurdin",
	}) // memasukan data temmplate
}

func TestFileTemplate(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(FileTemplate),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestFileTemplate2(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(FileTemplate2),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
