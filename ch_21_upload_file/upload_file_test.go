package ch_21_upload_file

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var myTemplates = template.Must(template.ParseGlob("./templates/*.html"))

func Form(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	file, fileHeader, err := request.FormFile("berkas_upload") // mengambil data dari name = "berkas_upload"
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename) // membuat folder resource dan menyimpan nama file dari fileHeader
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file) // mengcopy file ke file destination
	if err != nil {
		panic(err)
	}

	// ! Mengambil data name = "name"
	name := request.PostFormValue("nama")
	myTemplates.ExecuteTemplate(writer, "success", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}
func TestUpload(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Form)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/yes.jpeg
var yesJpeg []byte

func TestUploadUnitTest(t *testing.T) {
	body := new(bytes.Buffer) // membuat binary file
	writer := multipart.NewWriter(body)
	writer.WriteField("nama", "Foto Kucing")

	file, err := writer.CreateFormFile("berkas_upload", "yesUpload.jpeg")
	if err != nil {
		panic(err)
	}

	file.Write(yesJpeg)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:5000", body)
	req.Header.Set("Content-type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)

	bodyResponse, err := io.ReadAll(rec.Result().Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bodyResponse))
}
