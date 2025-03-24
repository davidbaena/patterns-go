package app

import (
	"encoding/json"
	"net/http"

	"bookstore/search/domain"
)

type BookController struct {
	bookSearchService *domain.BookSearchService
}

func NewBookController(bookSearchService *domain.BookSearchService) *BookController {
	return &BookController{bookSearchService: bookSearchService}
}

func (c *BookController) SearchByTitle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	books, err := c.bookSearchService.SearchByTitle(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) SearchByAuthor(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("authorID")
	books, err := c.bookSearchService.SearchByAuthor(authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func (c *BookController) SearchByGenre(w http.ResponseWriter, r *http.Request) {
	genreID := r.URL.Query().Get("genreID")
	books, err := c.bookSearchService.SearchByGenre(genreID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}
