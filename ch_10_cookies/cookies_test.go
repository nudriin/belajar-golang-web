package ch_10_cookies

import (
	"fmt"
	"net/http"
	"testing"
)

// Membuat response cookie
func SetCookiesHandler(writer http.ResponseWriter, request *http.Request) {
	// username := request.URL.Query().Get("username")
	// if username == "" {
	// 	writer.WriteHeader(400)
	// 	fmt.Fprint(writer, "username must provide")
	// } else {
	writer.WriteHeader(200)
	cookie := new(http.Cookie) // membuat cookie dengan function new
	cookie.Name = "X-cookie"   // set nama cookie
	cookie.Value = "Nurdin"    //  memberikan value cookie
	cookie.Path = "/"          // aktifnya di semua url

	http.SetCookie(writer, cookie) // ! Set cookies menggunakan set cookies
	fmt.Fprint(writer, "Success create cookies")
	// }
}

// Mengambil request cookie
func GetCookiesHandler(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("username")
	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookies(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookiesHandler)
	mux.HandleFunc("/get-cookie", GetCookiesHandler)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
