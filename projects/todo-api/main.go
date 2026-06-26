package main

import (
    "fmt"
    "log"
    "net/http"

    "todo-api/internal/handler"
    "todo-api/internal/store"
)

func main() {
    // Initialize the in-memory store
    s := store.New()

    // Create the handler with the store dependency
    h := handler.New(s)

    // Register the route handler
    // All requests to /todos/ are handled by HandleTodos
    http.HandleFunc("/todos", h.HandleTodos)
    http.HandleFunc("/todos/", h.HandleTodos)

    port := ":8080"
    fmt.Printf("Server running on http://localhost%s\n", port)
    fmt.Println("Endpoints:")
    fmt.Println("  GET    /todos        - List all todos")
    fmt.Println("  POST   /todos        - Create a todo")
    fmt.Println("  GET    /todos/{id}   - Get a todo by ID")
    fmt.Println("  PUT    /todos/{id}   - Update a todo")
    fmt.Println("  DELETE /todos/{id}   - Delete a todo")

    // Start the HTTP server
    // nil uses the DefaultServeMux (the one we registered handlers with)
    log.Fatal(http.ListenAndServe(port, nil))
}
