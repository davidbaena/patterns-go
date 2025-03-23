# ADR-002: Use of Hexagonal Architecture in Bookstore Project

## Context
The Bookstore project aims to manage a collection of books, authors, and customer orders. To ensure a clean separation of concerns and improve maintainability, we have decided to adopt Hexagonal Architecture (also known as Ports and Adapters).

## Decision
We will use Hexagonal Architecture to structure the Bookstore project. This involves defining the core domain logic in a central layer, with separate layers for input (driving) and output (driven) adapters.

### Core Concepts
1. **Domain**: The central part of the application containing the business logic.
2. **Ports**: Interfaces that define the entry points to the domain logic.
3. **Infrastructure**: Implementations of the ports, handling the interaction with external systems (e.g., databases, web services).

### Implementation
1. **Domain Layer**: Contains the core business logic, including entities, value objects, and domain services.
2. **Application Layer**: Contains use cases and application services that orchestrate the domain logic.
3. **Infrastructure Layer**: Contains the implementations of the ports, such as repositories, controllers, and external service clients.

## Consequences
- **Pros**:
    - Clear separation of concerns.
    - Improved testability and maintainability.
    - Flexibility to change external systems without affecting the core domain logic.

- **Cons**:
    - Initial complexity in setting up the project structure.
    - Requires a deep understanding of Hexagonal Architecture principles.

## Status
Accepted

## Date
24-03-2025

## Author
David Baena