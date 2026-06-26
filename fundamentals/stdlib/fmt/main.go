package main

import "fmt"

func main() {
    // Print and Println: basic output
    fmt.Print("Hello, ") // no newline at the end
    fmt.Println("world!") // adds a newline at the end

    name := "Lucas"
    age := 25

    // Printf: formatted output with verbs
    // %s = string, %d = integer, %f = float, %T = type, %v = default format
    fmt.Printf("Name: %s, Age: %d\n", name, age)
    fmt.Printf("Type of name: %T\n", name)
    fmt.Printf("Value (default format): %v\n", age)

    // Sprintf: returns the formatted string instead of printing it
    greeting := fmt.Sprintf("Hey, %s! You are %d years old.", name, age)
    fmt.Println(greeting)

    // Precision control with floats
    pi := 3.14159265359
    fmt.Printf("Pi with 2 decimal places: %.2f\n", pi)
    fmt.Printf("Pi with 4 decimal places: %.4f\n", pi)

    // Width and alignment
    fmt.Printf("|%10s|%-10s|\n", "right", "left") // positive = right align, negative = left align
    fmt.Printf("|%05d|\n", 42) // pad with zeros

    // Multiple values and types
    items := []string{"apple", "banana", "cherry"}
    count := 3
    price := 2.50
    fmt.Printf("We have %d items: %v, each costs $%.2f\n", count, items, price)

    // Errorf: creates an error with a formatted message (useful for error handling)
    err := fmt.Errorf("invalid age: %d (must be >= 0)", -5)
    fmt.Println("Error:", err)

    // Scanning user input
    var input string
    fmt.Print("Type something: ")
    fmt.Scanln(&input)
    fmt.Println("You typed:", input)
}
