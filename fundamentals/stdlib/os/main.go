package main

import (
    "fmt"
    "os"
)

func main() {
    // ----- Environment variables

    // Get: reads an env var (returns empty string if not set)
    home := os.Getenv("HOME")
    fmt.Println("HOME:", home)

    // LookupEnv: checks if an env var exists (returns value and a boolean)
    path, exists := os.LookupEnv("PATH")
    if exists {
        fmt.Println("PATH exists, length:", len(path))
    }

    // Setenv: sets an env var for the current process
    os.Setenv("MY_APP_ENV", "development")
    fmt.Println("MY_APP_ENV:", os.Getenv("MY_APP_ENV"))

    // ----- Command-line arguments
    // os.Args is a slice of strings with the program name and all arguments
    // Example: running "go run main.go hello world" gives:
    // os.Args[0] = "/tmp/go-build.../main"
    // os.Args[1] = "hello"
    // os.Args[2] = "world"
    fmt.Println("Args:", os.Args)
    fmt.Println("Program name:", os.Args[0])
    if len(os.Args) > 1 {
        fmt.Println("First argument:", os.Args[1])
    }

    // ----- File operations

    // WriteFile: creates or overwrites a file with the given content
    data := []byte("Hello from Go!\nThis is a test file.\n")
    err := os.WriteFile("test.txt", data, 0644)
    if err != nil {
        fmt.Println("WriteFile error:", err)
        return
    }
    fmt.Println("File 'test.txt' created!")

    // ReadFile: reads the entire file into memory
    content, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("ReadFile error:", err)
        return
    }
    fmt.Println("File content:", string(content))

    // ----- File info

    // Stat: gets file metadata without opening it
    info, err := os.Stat("test.txt")
    if err != nil {
        fmt.Println("Stat error:", err)
        return
    }
    fmt.Println("File name:", info.Name())
    fmt.Println("File size:", info.Size(), "bytes")
    fmt.Println("Is directory:", info.IsDir())
    fmt.Println("Modified at:", info.ModTime())

    // ----- Check if file exists
    // os.Stat returns an error if the file doesn't exist
    _, err = os.Stat("nonexistent.txt")
    if os.IsNotExist(err) {
        fmt.Println("File 'nonexistent.txt' does not exist")
    }

    // ----- Remove file
    err = os.Remove("test.txt")
    if err != nil {
        fmt.Println("Remove error:", err)
        return
    }
    fmt.Println("File 'test.txt' removed!")

    // ----- Working directory
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Println("Getwd error:", err)
        return
    }
    fmt.Println("Current directory:", cwd)

    // ----- Exit code (uncomment to use)
    // os.Exit(0) // exits with code 0 (success)
    // os.Exit(1) // exits with code 1 (error)
}
