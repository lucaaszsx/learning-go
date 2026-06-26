package main

import (
    "fmt"
    "strings"
)

// ----- Functions to be tested

// Add returns the sum of two integers
func Add(a, b int) int {
    return a + b
}

// Divide returns the result of dividing a by b, and an error if b is zero
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// IsPalindrome checks if a string is a palindrome (case-insensitive, ignores spaces)
func IsPalindrome(s string) bool {
    s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

// Factorial returns the factorial of n (n!)
func Factorial(n int) int {
    if n < 0 {
        return 0
    }
    if n <= 1 {
        return 1
    }
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    // ----- How testing works in Go
    // 1. Create a file ending in _test.go (e.g., main_test.go)
    // 2. Import the "testing" package
    // 3. Write functions starting with "Test" + name
    // 4. Run with: go test ./...
    // 5. Run verbose: go test -v ./...
    // 6. Run specific test: go test -run TestAdd ./...

    fmt.Println("Testing in Go - Examples of what tests look like:")
    fmt.Println()

    // ----- Test function structure
    // func TestAdd(t *testing.T) {
    //     result := Add(2, 3)
    //     if result != 5 {
    //         t.Errorf("Add(2, 3) = %d; want 5", result)
    //     }
    // }

    fmt.Println(`// Basic test:
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}`)

    // ----- Table-driven tests: test multiple cases in a loop
    fmt.Println(`
// Table-driven test:
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {100, 200, 300},
    }
    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}`)

    // ----- Test with error handling
    fmt.Println(`
// Test with error:
func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if result != 5.0 {
        t.Errorf("Divide(10, 2) = %f; want 5.0", result)
    }

    _, err = Divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}`)

    // ----- Subtests: organize related tests
    fmt.Println(`
// Subtests:
func TestIsPalindrome(t *testing.T) {
    tests := map[string]struct {
        input string
        want  bool
    }{
        "simple":       {"racecar", true},
        "with spaces":  {"race car", true},
        "mixed case":   {"RaceCar", true},
        "not palindrome": {"hello", false},
    }
    for name, tt := range tests {
        t.Run(name, func(t *testing.T) {
            got := IsPalindrome(tt.input)
            if got != tt.want {
                t.Errorf("IsPalindrome(%q) = %v; want %v", tt.input, got, tt.want)
            }
        })
    }
}`)

    // ----- Benchmarks: measure performance
    fmt.Println(`
// Benchmark:
func BenchmarkFactorial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Factorial(20)
    }
}
// Run with: go test -bench=. -benchmem`)

    // ----- Test with setup/teardown
    fmt.Println(`
// TestMain: setup/teardown for all tests in a package
func TestMain(m *testing.M) {
    // setup: runs before all tests
    fmt.Println("Setting up tests...")
    
    code := m.Run() // run all tests
    
    // teardown: runs after all tests
    fmt.Println("Tearing down tests...")
    os.Exit(code)
}`)

    // ----- Running tests
    fmt.Println("\nUseful commands:")
    fmt.Println("  go test ./...           Run all tests")
    fmt.Println("  go test -v ./...        Run tests with verbose output")
    fmt.Println("  go test -run TestAdd    Run specific test")
    fmt.Println("  go test -bench=.        Run benchmarks")
    fmt.Println("  go test -cover          Show test coverage")
    fmt.Println("  go test -coverprofile=c.out && go tool cover -html=c.out")
}
