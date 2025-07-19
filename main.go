package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port = ":5178"

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/welcome", basicHandler)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Println("Server is running on http://localhost" + port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error listening to server:", err)
		return
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome!")
	w.Write([]byte("Welcome!"))
}
