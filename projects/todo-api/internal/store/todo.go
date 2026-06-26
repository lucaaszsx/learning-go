package store

import (
    "fmt"
    "sync"
    "time"

    "todo-api/internal/model"
)

// TodoStore manages todo items in memory
// Uses a map for fast lookups and a RWMutex to safely handle concurrent access
type TodoStore struct {
    mu      sync.RWMutex
    todos   map[string]model.Todo
    counter int
}

// New creates a new empty TodoStore
func New() *TodoStore {
    return &TodoStore{
        todos: make(map[string]model.Todo),
    }
}

// GetAll returns all todos as a slice
func (s *TodoStore) GetAll() []model.Todo {
    s.mu.RLock()         // read lock: multiple readers allowed
    defer s.mu.RUnlock()

    result := make([]model.Todo, 0, len(s.todos))
    for _, todo := range s.todos {
        result = append(result, todo)
    }
    return result
}

// GetByID returns a single todo by its ID
// Second return value indicates if the todo was found
func (s *TodoStore) GetByID(id string) (model.Todo, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()

    todo, ok := s.todos[id]
    return todo, ok
}

// Create adds a new todo and returns it with the generated ID and timestamp
func (s *TodoStore) Create(todo model.Todo) model.Todo {
    s.mu.Lock()          // write lock: exclusive access
    defer s.mu.Unlock()

    s.counter++
    todo.ID = fmt.Sprintf("%d", s.counter)
    todo.CreatedAt = time.Now()
    todo.Completed = false

    s.todos[todo.ID] = todo
    return todo
}

// Update modifies an existing todo by ID
// Returns the updated todo and a boolean indicating if it was found
func (s *TodoStore) Update(id string, update model.Todo) (model.Todo, bool) {
    s.mu.Lock()
    defer s.mu.Unlock()

    existing, ok := s.todos[id]
    if !ok {
        return model.Todo{}, false
    }

    // Only update fields that were provided
    if update.Title != "" {
        existing.Title = update.Title
    }
    existing.Completed = update.Completed

    s.todos[id] = existing
    return existing, true
}

// Delete removes a todo by ID
// Returns true if the todo was found and deleted
func (s *TodoStore) Delete(id string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, ok := s.todos[id]
    if !ok {
        return false
    }

    delete(s.todos, id)
    return true
}
