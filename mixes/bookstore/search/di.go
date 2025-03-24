package search

import (
	"bookstore/search/app"
	"bookstore/search/domain"
	"bookstore/search/infrastructure"
)

func InitializeBookstore() *app.BookController {
	// Initialize repositories
	bookRepo := infrastructure.NewInMemoryBookRepository()
	//authorRepo := NewInMemoryAuthorRepository()
	//genreRepo := NewInMemoryGenreRepository()

	// Initialize use cases
	searchByTitle := domain.NewSearchBooksByTitle(bookRepo)
	searchByAuthor := domain.NewSearchBooksByAuthor(bookRepo)
	searchByGenre := domain.NewSearchBooksByGenre(bookRepo)

	// Initialize services
	bookSearchService := domain.NewBookSearchService(searchByTitle, searchByAuthor, searchByGenre)

	// Initialize controllers
	bookController := app.NewBookController(bookSearchService)

	return bookController
}
