package handler

import (
    "encoding/json"
    "net/http"
    "strings"

    "todo-api/internal/model"
    "todo-api/internal/store"
)

// Handler holds the store dependency and provides HTTP handlers for the Todo API
type Handler struct {
    store *store.TodoStore
}

// New creates a new Handler with the given store
func New(s *store.TodoStore) *Handler {
    return &Handler{store: s}
}

// HandleTodos routes requests to /todos based on the HTTP method
func (h *Handler) HandleTodos(w http.ResponseWriter, r *http.Request) {
    // Extract the ID from the URL path (e.g., /todos/1 → "1")
    id := strings.TrimPrefix(r.URL.Path, "/todos")

    // If there's no ID, the request is for the collection
    if id == "" || id == "/" {
        switch r.Method {
        case http.MethodGet:
            h.list(w, r)
        case http.MethodPost:
            h.create(w, r)
        default:
            respondError(w, http.StatusMethodNotAllowed, "method not allowed")
        }
        return
    }

    // Remove leading slash from the ID
    id = strings.TrimPrefix(id, "/")

    // Route to the appropriate handler based on method
    switch r.Method {
    case http.MethodGet:
        h.getByID(w, r, id)
    case http.MethodPut:
        h.update(w, r, id)
    case http.MethodDelete:
        h.delete(w, r, id)
    default:
        respondError(w, http.StatusMethodNotAllowed, "method not allowed")
    }
}

// list handles GET /todos - returns all todos
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
    todos := h.store.GetAll()
    respondJSON(w, http.StatusOK, todos)
}

// create handles POST /todos - creates a new todo
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
    var todo model.Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        respondError(w, http.StatusBadRequest, "invalid JSON body")
        return
    }

    if todo.Title == "" {
        respondError(w, http.StatusBadRequest, "title is required")
        return
    }

    created := h.store.Create(todo)
    respondJSON(w, http.StatusCreated, created)
}

// getByID handles GET /todos/{id} - returns a single todo
func (h *Handler) getByID(w http.ResponseWriter, r *http.Request, id string) {
    todo, ok := h.store.GetByID(id)
    if !ok {
        respondError(w, http.StatusNotFound, "todo not found")
        return
    }

    respondJSON(w, http.StatusOK, todo)
}

// update handles PUT /todos/{id} - updates an existing todo
func (h *Handler) update(w http.ResponseWriter, r *http.Request, id string) {
    var update model.Todo
    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        respondError(w, http.StatusBadRequest, "invalid JSON body")
        return
    }

    updated, ok := h.store.Update(id, update)
    if !ok {
        respondError(w, http.StatusNotFound, "todo not found")
        return
    }

    respondJSON(w, http.StatusOK, updated)
}

// delete handles DELETE /todos/{id} - removes a todo
func (h *Handler) delete(w http.ResponseWriter, r *http.Request, id string) {
    if !h.store.Delete(id) {
        respondError(w, http.StatusNotFound, "todo not found")
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// respondJSON writes a JSON response with the given status code
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

// respondError writes a JSON error response
func respondError(w http.ResponseWriter, status int, message string) {
    respondJSON(w, status, map[string]string{"error": message})
}
