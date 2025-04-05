# ADR-001: Use of Domain-Driven Design (DDD) in Bookstore Project

## Context
The Bookstore project aims to manage a collection of books, authors, and customer orders. To handle the complexity and ensure maintainability, we have decided to adopt Domain-Driven Design (DDD) principles.

## Decision
We will use DDD to structure the Bookstore project. This involves defining the core domains, subdomains, and bounded contexts. The project will be divided into several modules, each representing a specific domain or subdomain.

### Core Concepts
1. **Domain**: The main area of focus, which in this case is the Bookstore.
2. **Subdomains**: Specific areas within the main domain, such as Inventory, Sales, and Customer Management.
3. **Bounded Contexts**: Logical boundaries within which a particular model is defined and applicable.

### Implementation
1. **Entities**: Represent core objects with a distinct identity, such as `Book`, `Author`, and `Customer`.
2. **Value Objects**: Represent attributes that do not have a distinct identity, such as `Address` and `ISBN`.
3. **Aggregates**: Clusters of entities and value objects that are treated as a single unit, such as `Order` (which includes `OrderItem`).
4. **Repositories**: Interfaces for accessing aggregates, ensuring separation of concerns between the domain and data access layers.
5. **Services**: Encapsulate domain logic that does not naturally fit within entities or value objects.

## Consequences
- **Pros**:
    - Improved code organization and maintainability.
    - Clear separation of concerns.
    - Enhanced scalability and flexibility.

- **Cons**:
    - Initial complexity in setting up the project structure.
    - Requires a deep understanding of DDD principles.

## Date
24-04-2025

## Status
Accepted

## Author
David Baena