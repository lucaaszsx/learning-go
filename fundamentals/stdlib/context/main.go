package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // ----- Background context: the root of all contexts
    // context.Background() is the top-level context, typically used in main()
    ctx := context.Background()
    fmt.Println("Background context:", ctx)

    // ----- Context with timeout: auto-cancel after a duration
    // Useful for operations that should not take too long
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // ALWAYS call cancel to release resources

    // Simulate a slow operation
    result, err := slowOperation(ctx)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

    // ----- Context with deadline: cancel at a specific time
    deadline := time.Now().Add(1 * time.Second)
    ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
    defer cancel2()

    select {
    case <-time.After(500 * time.Millisecond):
        fmt.Println("Operation completed in time")
    case <-ctx2.Done():
        fmt.Println("Deadline exceeded:", ctx2.Err())
    }

    // ----- Context with value: pass data through the call chain
    // Use for request-scoped data (request IDs, user info, etc.)
    // Do NOT use for function parameters or optional arguments
    ctx3 := context.WithValue(context.Background(), "requestID", "abc-123")
    processRequest(ctx3)

    // ----- Context cancellation: manual cancel
    ctx4, cancel4 := context.WithCancel(context.Background())

    go func() {
        time.Sleep(500 * time.Millisecond)
        cancel4() // manually cancel after 500ms
    }()

    select {
    case <-time.After(2 * time.Second):
        fmt.Println("This won't print")
    case <-ctx4.Done():
        fmt.Println("Cancelled:", ctx4.Err()) // context canceled
    }

    // ----- Practical example: HTTP request with timeout
    // In real code, you'd pass ctx to http.NewRequestWithContext()
    // req, _ := http.NewRequestWithContext(ctx, "GET", "https://api.example.com", nil)
    // resp, err := client.Do(req)

    // ----- Checking context deadline
    ctx5, _ := context.WithTimeout(context.Background(), 1*time.Second)
    if deadline, ok := ctx5.Deadline(); ok {
        fmt.Println("Deadline:", deadline.Format("15:04:05"))
    }
}

// slowOperation simulates an operation that respects context cancellation
func slowOperation(ctx context.Context) (string, error) {
    select {
    case <-time.After(1 * time.Second): // simulates work
        return "operation completed", nil
    case <-ctx.Done(): // context was cancelled or timed out
        return "", ctx.Err()
    }
}

// processRequest demonstrates reading values from context
func processRequest(ctx context.Context) {
    // ctx.Value returns the value associated with the key
    if reqID, ok := ctx.Value("requestID").(string); ok {
        fmt.Println("Processing request:", reqID)
    }
}
