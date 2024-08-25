package ch_17_template_function

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + " my name is " + myPage.Name
}

func TemplateFunctionHandler(writer http.ResponseWriter, req *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Hishasy"}}`)) // memanggil function say hello yang ada pada struct

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Nurdin",
	})
}

func GlobalFuncs(writer http.ResponseWriter, req *http.Request) {
	t := template.New("GLOBALS")
	// ! membuat function global
	t = t.Funcs(map[string]any{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name}}`))
	t.ExecuteTemplate(writer, "GLOBALS", map[string]any{
		"Name": "Nurdin",
	})
}

func GlobalFuncsPipeline(writer http.ResponseWriter, req *http.Request) {
	t := template.New("GLOBALS")
	// ! membuat function global
	t = t.Funcs(map[string]any{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	// ! hasil dari sayHello aka diteruskan ke func upper
	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))
	t.ExecuteTemplate(writer, "GLOBALS", map[string]any{
		"Name": "Nurdin",
	})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	TemplateFunctionHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestGlobalFuncs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	GlobalFuncs(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
func TestGlobalFuncsPipeline(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	GlobalFuncsPipeline(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
