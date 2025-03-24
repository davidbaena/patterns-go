package main

import (
	"log"
	"net/http"

	"bookstore/search"
)

func main() {
	// Initialize the bookstore application
	bookController := search.InitializeBookstore()

	// Define API endpoints
	http.HandleFunc("/api/books/search/title", bookController.SearchByTitle)
	http.HandleFunc("/api/books/search/author", bookController.SearchByAuthor)
	http.HandleFunc("/api/books/search/genre", bookController.SearchByGenre)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
