package main

import (
    "errors"
    "fmt"
)

func main() {
    // Creating errors with errors.New
    err := errors.New("something went wrong")
    fmt.Println("Error:", err)

    // Using fmt.Errorf for formatted error messages
    name := "Lucas"
    age := -5
    err = fmt.Errorf("invalid age for %s: %d", name, age)
    fmt.Println("Formatted error:", err)

    // Comparing errors with errors.Is
    ErrNotFound := errors.New("not found")
    ErrPermission := errors.New("permission denied")

    result, err := findUser("unknown")
    if err != nil {
        // errors.Is checks if the error (or any wrapped error) matches the target
        if errors.Is(err, ErrNotFound) {
            fmt.Println("User not found!")
        } else if errors.Is(err, ErrPermission) {
            fmt.Println("Permission denied!")
        }
    }
    fmt.Println("Result:", result)

    // Wrapping errors: add context while preserving the original error
    err = openConfig("config.json")
    if err != nil {
        fmt.Println("Wrapped error:", err)

        // errors.Is still works through the wrapping chain
        if errors.Is(err, ErrNotFound) {
            fmt.Println("Config file not found (detected through wrapping)")
        }
    }

    // errors.As: extract a specific error type
    err = validateAge(-5)
    var valErr *ValidationError
    if errors.As(err, &valErr) {
        fmt.Printf("Validation failed: field=%s, message=%s\n", valErr.Field, valErr.Message)
    }
}

// ErrNotFound is a sentinel error: a package-level variable that represents a specific error condition
var ErrNotFound = errors.New("not found")

// findUser simulates a function that might return different errors
func findUser(name string) (string, error) {
    if name == "unknown" {
        return "", ErrNotFound
    }
    return name, nil
}

// openConfig wraps an error with additional context using fmt.Errorf and %w
func openConfig(path string) error {
    _, err := findUser(path)
    if err != nil {
        // %w wraps the original error, making it accessible via errors.Is and errors.As
        return fmt.Errorf("open config %s: %w", path, err)
    }
    return nil
}

// ValidationError is a custom error type with additional context
type ValidationError struct {
    Field   string
    Message string
}

// implementing the error interface (any type with Error() string method is an error)
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// validateAge returns a custom error type
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "must be a positive number",
        }
    }
    return nil
}
