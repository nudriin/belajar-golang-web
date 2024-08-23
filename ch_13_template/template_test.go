package ch_13_template

import (
	"html/template"
	"net/http"
	"testing"
)

func SimpleTemplate(writer http.ResponseWriter, request *http.Request) {
	templateText := "<html><body><h1>{{.}}</h1></body></html>"
	// ! membuat template
	templates := template.New("NURDIN")

	// ! PArsing template
	t, err := templates.Parse(templateText)
	if err != nil {
		panic(err)
	}

	name := request.URL.Query().Get("name")

	t.ExecuteTemplate(writer, "NURDIN", "Hello "+name) // execute template dan mengrimkannya sebagi response
}

func FileTemplate(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "template.html", "Hello")
}

func FileTemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.html") // parsing semua yang ada pada directory templates
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "template.html", "Hello")
}

func TestTemplate(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(SimpleTemplate),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
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

func TestFileTemplateDirectpry(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: http.HandlerFunc(FileTemplateDirectory),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
