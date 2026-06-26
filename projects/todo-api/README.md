# Todo API

A simple REST API for managing todo items, built with Go's standard library.

## Structure

```
todo-api/
├── main.go                  # HTTP server entry point
├── go.mod                   # Go module definition
└── internal/
    ├── model/todo.go        # Todo data structure
    ├── store/todo.go        # In-memory storage with RWMutex
    └── handler/todo.go      # HTTP request handlers
```

## How it works

- **model**: defines the `Todo` struct with JSON tags
- **store**: manages a `map[string]Todo` protected by `sync.RWMutex` for safe concurrent access
- **handler**: routes HTTP requests to the appropriate CRUD operations and returns JSON responses

## Endpoints

| Method | Path | Description | Request Body |
|--------|------|-------------|--------------|
| GET | `/todos` | List all todos | - |
| POST | `/todos` | Create a new todo | `{"title": "..."}` |
| GET | `/todos/{id}` | Get a todo by ID | - |
| PUT | `/todos/{id}` | Update a todo | `{"title": "...", "completed": true}` |
| DELETE | `/todos/{id}` | Delete a todo | - |

## Running

```bash
cd projects/todo-api
go run main.go
```

The server starts on `http://localhost:8080`.

## Testing with curl

```bash
# Create a todo
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go"}'

# List all todos
curl http://localhost:8080/todos

# Get a specific todo
curl http://localhost:8080/todos/1

# Update a todo
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go APIs","completed":true}'

# Delete a todo
curl -X DELETE http://localhost:8080/todos/1
```

## Error responses

All errors are returned as JSON:

```json
{
  "error": "todo not found"
}
```

## Standard library packages used

- `net/http` - HTTP server and routing
- `encoding/json` - JSON encoding/decoding
- `sync` - RWMutex for concurrent access
- `time` - timestamps on todo items
- `strings` - URL path parsing
