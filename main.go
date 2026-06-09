package main

import (
	"log"
	"net/http"

	"github.com/demo/api/handlers"
	"github.com/demo/api/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/users", handlers.UsersHandler)
	mux.HandleFunc("/users/", handlers.UserByIDHandler)
	mux.HandleFunc("/products", handlers.ProductsHandler)

	handler := middleware.Logger(mux)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
