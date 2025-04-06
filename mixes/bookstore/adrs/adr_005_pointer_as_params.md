# ADR-005: Use of Pointers as Parameters for Mutating Functions in Bookstore Project

## Context

The Bookstore project aims to manage a collection of books, authors, and customer orders. To ensure clarity and maintainability, we have decided to adopt a convention for using pointers as parameters in functions. Specifically, we will use pointers only when a function mutates the attributes of a parameter. This approach will help developers easily identify which functions modify the state of an object and which do not.

## Decision

We will use pointers as parameters in functions only when the function mutates the attributes of the parameter. This approach will help developers easily identify which functions modify the state of an object and maintain consistency within the codebase.

### Implementation

1. **Non-Mutating Functions**: Use value parameters for functions that do not modify the parameter's attributes.
2. **Mutating Functions**: Use pointer parameters for functions that modify the parameter's attributes.

### Example

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

type Library struct {
    Books []Book
}

// Non-mutating function
func GetBookInfo(b Book) string {
    return b.Title + " by " + b.Author
}

// Mutating function
func UpdateLibraryBookDetails(lib *Library, index int, title string, author string, pages int) {
    if index >= 0 && index < len(lib.Books) {
        lib.Books[index].Title = title
        lib.Books[index].Author = author
        lib.Books[index].Pages = pages
    }
}
```