package main

import (
    "fmt"
    "net/http"
    "time"
)

// ----- HTTP Client: making requests to external APIs

func main() {
    // http.Get: simplest way to make a GET request
    // Returns a Response and an error
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        fmt.Println("GET error:", err)
        return
    }
    // ALWAYS close the response body when done
    defer resp.Body.Close()

    // Response fields
    fmt.Println("Status:", resp.Status)           // "200 OK"
    fmt.Println("StatusCode:", resp.StatusCode)    // 200
    fmt.Println("Content-Length:", resp.ContentLength)

    // Reading the response body
    // body, err := io.ReadAll(resp.Body)
    // if err != nil {
    //     fmt.Println("Read body error:", err)
    //     return
    // }
    // fmt.Println("Body:", string(body))

    // ----- Custom HTTP Client with configuration
    // Use when you need timeouts, redirects, or TLS settings
    client := &http.Client{
        Timeout: 10 * time.Second,
    }

    // http.NewRequest: create a custom request (for POST, PUT, DELETE, headers, etc.)
    req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
    if err != nil {
        fmt.Println("NewRequest error:", err)
        return
    }

    // Adding custom headers
    req.Header.Set("User-Agent", "learning-go/1.0")
    req.Header.Set("Accept", "application/json")

    resp2, err := client.Do(req)
    if err != nil {
        fmt.Println("Client.Do error:", err)
        return
    }
    defer resp2.Body.Close()
    fmt.Println("Custom request status:", resp2.Status)

    // ----- HTTP Server: creating a simple web server

    // http.HandleFunc: register a handler for a URL pattern
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/api/user", userHandler)

    fmt.Println("Server starting on :8080...")
    fmt.Println("Try: http://localhost:8080/")
    fmt.Println("Try: http://localhost:8080/hello?name=Lucas")
    fmt.Println("Try: http://localhost:8080/api/user")

    // http.ListenAndServe: start the server
    // First argument: address to listen on (":8080" = all interfaces, port 8080)
    // Second argument: handler (nil uses DefaultServeMux)
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Server error:", err)
    }
}

// homeHandler handles requests to "/"
// w = response writer (send data back to client)
// r = request object (contains all request data)
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Go HTTP server!")
}

// helloHandler demonstrates query parameters
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // r.URL.Query() parses the query string
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "World"
    }
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Hello, %s!", name)
}

// userHandler demonstrates returning JSON
func userHandler(w http.ResponseWriter, r *http.Request) {
    // Only allow GET method
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"name":"Lucas","age":25,"active":true}`)
}
