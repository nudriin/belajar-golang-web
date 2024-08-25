package ch_15_template_action

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func ActionHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/template.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "template.html", map[string]any{
		"Name": req.URL.Query().Get("name"),
	})
}

func ComparatorHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/comparator.html")
	if err != nil {
		panic(err)
	}

	value, err := strconv.Atoi(req.URL.Query().Get("value"))
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "comparator.html", map[string]any{
		"NilaiAkhir": value,
	})
}

func RangeHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/range.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "range.html", map[string]any{
		"Hobbies": []string{
			"Menyanyi",
			"Berlari",
			"Belajar",
			"Bersepeda",
		},
	})
}

type Address struct {
	Street string
	City   string
}

func WithHandler(writer http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("./templates/with.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "with.html", map[string]any{
		"Name": "Nurdin",
		"Address": Address{
			Street: "JL. Pangeran Samuder",
			City:   "Palangka Raya",
		},
	})
}

func TestTemplateAction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000?name=Nurdin", nil)
	rec := httptest.NewRecorder()

	ActionHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	req = httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec = httptest.NewRecorder()

	ActionHandler(rec, req)

	body, err = io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestComparator(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000?value=80", nil)
	rec := httptest.NewRecorder()

	ComparatorHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
func TestRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	RangeHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}

func TestWith(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	rec := httptest.NewRecorder()

	WithHandler(rec, req)

	body, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
