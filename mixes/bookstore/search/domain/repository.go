package domain

// BookRepository defines the interface for book repository operations.
type BookRepository interface {
	FindByID(id string) (*Book, error)
	FindByTitle(title string) ([]*Book, error)
	FindByAuthor(authorID string) ([]*Book, error)
	FindByGenre(genreID string) ([]*Book, error)
	Save(book *Book) error
	Delete(id string) error
}

// AuthorRepository defines the interface for author repository operations.
type AuthorRepository interface {
	FindByID(id string) (*Author, error)
	FindByName(name string) ([]*Author, error)
	Save(author *Author) error
	Delete(id string) error
}

// GenreRepository defines the interface for genre repository operations.
type GenreRepository interface {
	FindByID(id string) (*Genre, error)
	FindByName(name string) ([]*Genre, error)
	Save(genre *Genre) error
	Delete(id string) error
}
