# ADR-004: Use of Pointers for Mutating Functions in Bookstore Project

## Context

The Bookstore project aims to manage a collection of books, authors, and customer orders. To ensure clarity and maintainability, we have decided to adopt a convention for using pointers in function receivers. Specifically, we will use pointers only when a function mutates the attributes of a struct. However, if a service contains at least one pointer receiver, all receivers in that service will be pointers for consistency.

## Decision

We will use pointers in function receivers only when the function mutates the attributes of the struct. If a service contains at least one pointer receiver, all receivers in that service will be pointers. This approach will help developers easily identify which functions modify the state of an object and maintain consistency within services.

### Implementation

1. **Non-Mutating Functions**: Use value receivers for functions that do not modify the struct's attributes, unless the service contains a mutating function.
2. **Mutating Functions**: Use pointer receivers for functions that modify the struct's attributes. If a service contains a mutating function, all functions in that service will use pointer receivers.

### Example

```go
type Book struct {
    Title  string
    Author string
}

// Service with no mutating functions
type BookService struct{}

func (s BookService) GetTitle(b Book) string {
    return b.Title
}

// Service with mutating functions
type BookServiceWithMutations struct{}

func (s *BookServiceWithMutations) GetTitle(b *Book) string {
    return b.Title
}

func (s *BookServiceWithMutations) SetTitle(b *Book, title string) {
    b.Title = title
}
```

## Consequences

- **Pros**:
    - Improved code clarity and maintainability.
    - Easier to identify functions that modify the state of an object.
    - Consistent use of pointers within services.

- **Cons**:
    - Requires discipline and consistency in following the convention.
    - Potential for confusion if the convention is not well-documented or understood.

## Date
06-04-2025

## Status
Accepted

## Author
David Baena
