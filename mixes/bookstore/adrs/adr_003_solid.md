# ADR: Use of SOLID Principles in Bookstore Project

## Context

The Bookstore project aims to manage a collection of books, authors, and customer orders. To ensure the project is maintainable, scalable, and easy to understand, we have decided to adopt the SOLID principles of object-oriented design.

## Decision

We will use the SOLID principles to guide the design and implementation of the Bookstore project. The SOLID principles are:

1. **Single Responsibility Principle (SRP)**: A class should have only one reason to change, meaning it should have only one job or responsibility.
2. **Open/Closed Principle (OCP)**: Software entities should be open for extension but closed for modification.
3. **Liskov Substitution Principle (LSP)**: Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.
4. **Interface Segregation Principle (ISP)**: Clients should not be forced to depend on interfaces they do not use.
5. **Dependency Inversion Principle (DIP)**: High-level modules should not depend on low-level modules. Both should depend on abstractions.

## Consequences

- **Pros**:
    - Improved code organization and maintainability.
    - Enhanced flexibility and scalability.
    - Easier to understand and modify the codebase.
    - Better separation of concerns.

- **Cons**:
    - Initial complexity in understanding and applying the principles.
    - Potential for over-engineering if not applied judiciously.

## Date
24-03-2025

## Status
Accepted

## Author
David Baena