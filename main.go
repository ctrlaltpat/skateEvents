package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var port = ":5178"

func main() {
	server := &http.Server{
		Addr:    port,
		Handler: http.HandlerFunc(basicHandler),
	}

	fmt.Println("Server is running on http://localhost" + port)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error listening to server:", err)
		return
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, World!")
	w.Write([]byte("Hello, World!"))
}
