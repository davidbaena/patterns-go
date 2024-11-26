package infrastructure

import (
	"errors"

	"bookstore/bookstore/domain"
)

// InMemoryAuthorRepository is an in-memory implementation of AuthorRepository.
type InMemoryAuthorRepository struct {
	authors map[string]*domain.Author
}

func NewInMemoryAuthorRepository() *InMemoryAuthorRepository {
	return &InMemoryAuthorRepository{authors: make(map[string]*domain.Author)}
}

func (r *InMemoryAuthorRepository) FindByID(id string) (*domain.Author, error) {
	if author, exists := r.authors[id]; exists {
		return author, nil
	}
	return nil, errors.New("author not found")
}

func (r *InMemoryAuthorRepository) FindByName(name string) ([]*domain.Author, error) {
	var result []*domain.Author
	for _, author := range r.authors {
		if author.Name == name {
			result = append(result, author)
		}
	}
	return result, nil
}

func (r *InMemoryAuthorRepository) Save(author *domain.Author) error {
	r.authors[author.ID] = author
	return nil
}

func (r *InMemoryAuthorRepository) Delete(id string) error {
	delete(r.authors, id)
	return nil
}
