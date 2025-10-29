# Islamic Application API

A Go-based REST API service built with the Gin framework that provides todo management functionality.

## Description

This project is a RESTful API service built with Go and Gin framework that implements a todo management system. It provides endpoints for creating, reading, and deleting todo items.

## Prerequisites

- Go 1.x or higher
- Git
- [Gin Framework](https://github.com/gin-gonic/gin)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/muradwalana/IslamicApplicatoinAPI.git
cd IslamicApplicatoinAPI
```

2. Install dependencies:
```bash
go mod tidy
```

## Running the Application

To start the server:
```bash
go run main.go
```

The server will start on `localhost:8080`

## API Endpoints

### Data Structure

The API uses the following Todo structure:
```go
type todo struct {
    ID   string `json:"id"`
    Task string `json:"task"`
    Done bool   `json:"done"`
}
```

### Available Endpoints

1. **Get All Todos**
   - Endpoint: `GET /todos-get`
   - Description: Retrieves all todos
   - Response: Array of todo objects
   ```bash
   curl http://localhost:8080/todos-get
   ```

2. **Get Todo by ID**
   - Endpoint: `GET /todos-get/:id`
   - Description: Retrieves a specific todo by ID
   - Response: Single todo object or 404 if not found
   ```bash
   curl http://localhost:8080/todos-get/1
   ```

3. **Create Todo**
   - Endpoint: `POST /todos-set`
   - Description: Creates a new todo
   - Request Body: Todo object
   ```bash
   curl -X POST http://localhost:8080/todos-set \
        -H "Content-Type: application/json" \
        -d '{"id":"4","task":"New Task","done":false}'
   ```

4. **Delete Todo**
   - Endpoint: `DELETE /todos-delete/:id`
   - Description: Deletes a todo by ID
   - Response: Success message or 404 if not found
   ```bash
   curl -X DELETE http://localhost:8080/todos-delete/1
   ```

## Sample Responses

### Get All Todos Response
```json
[
    {
        "id": "1",
        "task": "Buy groceries",
        "done": false
    },
    {
        "id": "2",
        "task": "Walk the dog",
        "done": true
    },
    {
        "id": "3",
        "task": "Read a book",
        "done": false
    }
]
```

### Error Response (404 Not Found)
```json
{
    "message": "Todo not found"
}
```

## Error Handling

The API implements basic error handling:
- Returns 404 status for not found resources
- Returns 201 status for successful creation
- Returns 200 status for successful operations

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your chosen license]

## Development Notes

- The API is built using the Gin framework for efficient routing and middleware support
- Uses in-memory storage for todos (can be extended to use a database)
- Implements RESTful conventions for endpoint design
- Provides JSON responses with proper status codes
