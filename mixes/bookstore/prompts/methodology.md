
### Step 1: Define the Task
- "Describe the task you want to implement, including the main goal, requirements, and constraints."
- "What are the expected outcomes of this task? Please provide specific details."
**Prompt Example:**
- "Define a task feature to allow users to search for books by genre. 
The main goal is to enhance the user experience by making it easier to find books of interest. 
Requirements include updating the book model to include genres, adding a search filter on the book listing page,
and ensuring the backend supports genre-based searches."

### Step 2: Break Down the Task
- "List the main components or sub-tasks needed to complete the task."
- "For each sub-task, describe what needs to be done and any dependencies it might have."
**Prompt Example:**
- "Sub-tasks for adding a book search by genre:
  1. Update the book model to include a genre attribute.
  2. Modify the database schema to store genres.
  3. Implement the backend logic to filter books by genre.
  4. Update the book listing page to include a genre filter.
  5. Write tests for the new functionality."

### Step 3: Design the Solution
**Prompt:**
- "Outline the architecture and design for the solution. What components and modules will you need?"
- "How will the components interact with each other? Provide a high-level overview."
- "Which design patterns or principles will you apply to ensure the solution is maintainable and scalable?"
**Prompt Example:**
- "Design for book search by genre:
  - **Components:** Book model, BookRepository, BookService, BookController, BookListingPage.
  - **Interaction:** The BookController will handle requests for genre-based searches, calling the BookService, which interacts with the BookRepository to fetch filtered results from the database.
  - **Design Patterns:** Use the Repository pattern for data access, and the Service pattern for business logic."

### Step 4: Implement the Code
- "Write the code for the first sub-task. Include any necessary classes, methods, and functions."
- "For each sub-task, provide the implementation details and explain how it integrates with the overall solution."
- "Ensure the code follows best practices and coding standards. What are those standards?"
**Prompt Copilot Example:**
- "Implement the book model update to include a genre attribute:
  - **Book Model:** Add a genre field to the Book model class.
  - **Migration:** Create a database migration to update the book table schema.
  - **Validation:** Add validation rules for the genre field to ensure data integrity.
  - **Testing:** Write unit tests to validate the model changes and migration."

### Step 5: Test the Implementation

- "Write unit tests for the code you implemented. What scenarios do you need to test?"
- "Describe how you will perform integration testing to ensure all components work together correctly."
- "What tools or frameworks will you use for testing? Provide examples of test cases."
**Prompt Example:**
- "Test the genre-based book search functionality:
  - **Unit Tests:** Verify the book model includes the genre field and validation rules.
  - **Integration Tests:** Test the interaction between the BookController, BookService, and BookRepository for genre-based searches.
  - **Tools:** Use Jest for unit testing and Supertest for API integration testing."

Example:
1) Given the requirement number 3 for the bookstore define the task
2) Break down the task into sub-tasks
3) Design the solution for the task
4) Implement the code for the first sub-task
5) Test the implementation of the code

### TRICKS
- I'm going to give you a task description but I don't want to implement directly. Wait until I give you the implementation steps.
- I want to extend the description to pass to GitHub Copilot to implement the code in Golang.the target is not to give me the code already implemented. I need to focus in the definition
- I want to enforce to use hexagonal Architecture and SOLID principles and DDD. I want to define the subtask 3 with that requisites.