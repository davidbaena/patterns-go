package infrastructure

import (
	"errors"

	"bookstore/bookstore/domain"
)

// InMemoryBookRepository is an in-memory implementation of BookRepository.
type InMemoryBookRepository struct {
	books map[string]*domain.Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	//populate the repository with some books
	books := map[string]*domain.Book{
		"1": &domain.Book{
			ID:    "1",
			Title: "The Great Gatsby",
		},
		"2": &domain.Book{
			ID:    "2",
			Title: "To Kill a Mockingbird",
		},
	}
	return &InMemoryBookRepository{books: books}
}

func (r *InMemoryBookRepository) FindByID(id string) (*domain.Book, error) {
	if book, exists := r.books[id]; exists {
		return book, nil
	}
	return nil, errors.New("book not found")
}

func (r *InMemoryBookRepository) FindByTitle(title string) ([]*domain.Book, error) {
	var result []*domain.Book
	for _, book := range r.books {
		if book.Title == title {
			result = append(result, book)
		}
	}
	return result, nil
}

func (r *InMemoryBookRepository) FindByAuthor(authorID string) ([]*domain.Book, error) {
	var result []*domain.Book
	for _, book := range r.books {
		if book.Author.ID == authorID {
			result = append(result, book)
		}
	}
	return result, nil
}

func (r *InMemoryBookRepository) FindByGenre(genreID string) ([]*domain.Book, error) {
	var result []*domain.Book
	for _, book := range r.books {
		if book.Genre.ID == genreID {
			result = append(result, book)
		}
	}
	return result, nil
}

func (r *InMemoryBookRepository) Save(book *domain.Book) error {
	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepository) Delete(id string) error {
	delete(r.books, id)
	return nil
}
