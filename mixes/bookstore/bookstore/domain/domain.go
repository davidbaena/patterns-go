package domain

// Book represents a book entity in the bookstore system.
type Book struct {
	ID     string
	Title  string
	Author Author
	Genre  Genre
	Price  float64
	Stock  int
}

// Author represents an author entity in the bookstore system.
type Author struct {
	ID   string
	Name string
}

// Genre represents a genre entity in the bookstore system.
type Genre struct {
	ID   string
	Name string
}
