package domain

type SearchBooksByTitle struct {
	bookRepo  BookRepository
	genreRepo GenreRepository
}

func NewSearchBooksByTitle(bookRepo BookRepository) *SearchBooksByTitle {
	return &SearchBooksByTitle{bookRepo: bookRepo}
}

func (uc *SearchBooksByTitle) Execute(title string) ([]*Book, error) {
	return uc.bookRepo.FindByTitle(title)
}

type SearchBooksByAuthor struct {
	bookRepo BookRepository
}

func NewSearchBooksByAuthor(bookRepo BookRepository) *SearchBooksByAuthor {
	return &SearchBooksByAuthor{bookRepo: bookRepo}
}

func (uc *SearchBooksByAuthor) Execute(authorID string) ([]*Book, error) {
	return uc.bookRepo.FindByAuthor(authorID)
}

type SearchBooksByGenre struct {
	bookRepo BookRepository
}

func NewSearchBooksByGenre(bookRepo BookRepository) *SearchBooksByGenre {
	return &SearchBooksByGenre{bookRepo: bookRepo}
}

func (uc *SearchBooksByGenre) Execute(genreID string) ([]*Book, error) {
	return uc.bookRepo.FindByGenre(genreID)
}
