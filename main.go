package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)
	server := &http.Server{
		Addr: ":5000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to servel", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}