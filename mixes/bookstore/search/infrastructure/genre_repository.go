package infrastructure

import (
	"errors"

	"bookstore/search/domain"
)

// InMemoryGenreRepository is an in-memory implementation of GenreRepository.
type InMemoryGenreRepository struct {
	genres map[string]*domain.Genre
}

func NewInMemoryGenreRepository() *InMemoryGenreRepository {
	return &InMemoryGenreRepository{genres: make(map[string]*domain.Genre)}
}

func (r *InMemoryGenreRepository) FindByID(id string) (*domain.Genre, error) {
	if genre, exists := r.genres[id]; exists {
		return genre, nil
	}
	return nil, errors.New("genre not found")
}

func (r *InMemoryGenreRepository) FindByName(name string) ([]*domain.Genre, error) {
	var result []*domain.Genre
	for _, genre := range r.genres {
		if genre.Name == name {
			result = append(result, genre)
		}
	}
	return result, nil
}

func (r *InMemoryGenreRepository) Save(genre *domain.Genre) error {
	r.genres[genre.ID] = genre
	return nil
}

func (r *InMemoryGenreRepository) Delete(id string) error {
	delete(r.genres, id)
	return nil
}
