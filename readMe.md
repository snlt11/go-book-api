# Go-Book-API

A RESTful API for a book store built in Go using [Gin](https://github.com/gin-gonic/gin) for routing, [GORM](https://gorm.io/) with SQLite for data persistence, and JWT for authentication.

## Features

- **User Authentication**
  - User registration with unique email validation and password hashing.
  - Login with JWT token generation.
- **Book CRUD**
  - Create, Read, Update, and Delete books.
  - Pagination and search filtering for listing books.
- **JWT Protected Routes**
  - Secure endpoints using JWT middleware.
- **Clean Project Structure**
  - Organized into packages: `config`, `controllers`, `middleware`, `models`, `routes`, `response`, and `utils`.

## Prerequisites

- [Go 1.18+](https://golang.org/dl/)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/snlt11/go-book-api.git
   cd go-book-api
   ```

2. **Initialize the Go module and download dependencies:**

   ```bash
   go mod tidy
   ```
   
## Project Structure

```
go-book-api/
├── config/                 # Database connection and configuration
│   └── db.go
├── controllers/            # HTTP handlers for authentication and books
│   ├── auth_controller.go
│   └── book_controller.go
├── middleware/             # JWT authentication middleware
│   └── jwt_middleware.go
├── models/                 # Database models
│   ├── book.go
│   ├── user.go
│   └── userResponse.go
├── routes/                 # API route registration
│   ├── auth_routes.go
│   └── book_routes.go
├── response/               # Response structs for clean output
│   └── userResponse.go
├── utils/                  # Utility functions (e.g., JWT generation)
│   └── token.go
├── main.go                 # Application entry point
└── go.mod                  # Go module file
```

## Running the Application

1. **Run the Server:**

   From the project root, run:

   ```bash
   go run main.go
   ```

   The server should start and listen on port `8080` (by default). You’ll see a log message similar to:

   ```
   [GIN-debug] Listening and serving HTTP on :8080
   ```

2. **API Endpoints:**

   - **Authentication:**
     - **Register:** `POST /auth/register`
       - **Body (JSON):**
         ```json
         {
           "username": "your_username",
           "email": "your_email@example.com",
           "password": "your_password"
         }
         ```
     - **Login:** `POST /auth/login`
       - **Body (JSON):**
         ```json
         {
           "email": "your_email@example.com",
           "password": "your_password"
         }
         ```
       - **Response:**
         ```json
         {
           "token": "<JWT token>"
         }
         ```

   - **Books (Protected, requires JWT in Authorization header):**
     - **List Books:** `GET /books?page=1&limit=10&search=keyword`
     - **Create Book:** `POST /books`
       - **Body (JSON):**
         ```json
         {
           "title": "The Great Gatsby",
           "author": "F. Scott Fitzgerald"
         }
         ```
     - **Update Book:** `PUT /books/{id}`
       - **Body (JSON):**
         ```json
         {
           "title": "The Great Gatsby - Updated",
           "author": "F. Scott Fitzgerald"
         }
         ```
     - **Delete Book:** `DELETE /books/{id}`

   **Note:** When making requests to protected routes, include the JWT token in the `Authorization` header:

   ```
   Authorization: <your_jwt_token>
   ```

## Testing with Postman

- **Register a User:**
  1. Send a `POST` request to `http://localhost:8080/auth/register` with the user JSON.
- **Login:**
  1. Send a `POST` request to `http://localhost:8080/auth/login` with the login JSON.
  2. Copy the returned JWT token.
- **Test Protected Routes:**
  1. For book operations (`POST`, `PUT`, `DELETE`, `GET /books`), add the JWT token in the header:
     - Key: `Authorization`
     - Value: `<your_jwt_token>`

## Database Migration

The app uses GORM's `AutoMigrate` function to automatically create/update the tables in the SQLite database (`books.db`). If you want to refresh migrations during development, you can either:

- Manually delete the `books.db` file and restart the app.
- Or modify the `ConnectDatabase` function in `config/db.go` to drop tables before migrating.

## Contributing

Feel free to open issues or submit pull requests to help improve this project.

## License

This project is licensed under the MIT License.
