package main

import (
	"log"
	"net/http"

	"bookstore/auth/app"
	"bookstore/auth/domain"
	"bookstore/auth/infrastructure"
	"bookstore/search"
)

func main() {
	// Initialize the bookstore application
	bookController := search.InitializeBookstore()

	// Initialize the auth domain
	userRepository := infrastructure.NewInMemoryUserRepository()
	passwordHasher := infrastructure.NewBcryptPasswordHasher()
	authService := domain.NewAuthService(userRepository, passwordHasher)
	authController := app.NewAuthController(authService)

	// Define book search API endpoints
	http.HandleFunc("/api/books/search/title", bookController.SearchByTitle)
	http.HandleFunc("/api/books/search/author", bookController.SearchByAuthor)
	http.HandleFunc("/api/books/search/genre", bookController.SearchByGenre)

	// Define auth API endpoints
	http.HandleFunc("/api/auth/register", authController.Register)
	http.HandleFunc("/api/auth/login", authController.Login)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
