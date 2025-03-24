package domain

type BookSearchService struct {
	searchByTitle  *SearchBooksByTitle
	searchByAuthor *SearchBooksByAuthor
	searchByGenre  *SearchBooksByGenre
}

func NewBookSearchService(
	searchByTitle *SearchBooksByTitle,
	searchByAuthor *SearchBooksByAuthor,
	searchByGenre *SearchBooksByGenre,
) *BookSearchService {
	return &BookSearchService{
		searchByTitle:  searchByTitle,
		searchByAuthor: searchByAuthor,
		searchByGenre:  searchByGenre,
	}
}

func (s *BookSearchService) SearchByTitle(title string) ([]*Book, error) {
	return s.searchByTitle.Execute(title)
}

func (s *BookSearchService) SearchByAuthor(authorID string) ([]*Book, error) {
	return s.searchByAuthor.Execute(authorID)
}

func (s *BookSearchService) SearchByGenre(genreID string) ([]*Book, error) {
	return s.searchByGenre.Execute(genreID)
}
