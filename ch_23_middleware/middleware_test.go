package ch_23_middleware

import (
	"fmt"
	"net/http"
	"testing"
)

type Middleware struct {
	Handler http.Handler
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After middleware")
}

type ErrorMiddleware struct {
	Handler http.Handler
}

func (errMiddleware *ErrorMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// * Handling error dengan errMiddleware
	defer func() {
		err := recover()
		fmt.Println("RECOVER :", err)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(writer, "ERROR :", err)
		}
	}()

	fmt.Println("Before errMiddleware")
	errMiddleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After errMiddleware")
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler execute")
		fmt.Fprint(writer, "This is handler")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic execute")
		panic("Some error hapens")
	})

	// * membuat middleware
	middleware := &Middleware{
		Handler: mux, // * menginisiasikan Handler dengan mux
	}

	// * Membuat error middleware yang ada di atas dari middleware
	errMiddleware := &ErrorMiddleware{
		Handler: middleware,
	}

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: errMiddleware, // * menggunakan middleware sebagai handler
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
