package bookstore

import (
	"bookstore/bookstore/app"
	"bookstore/bookstore/domain"
	"bookstore/bookstore/infrastructure"
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
