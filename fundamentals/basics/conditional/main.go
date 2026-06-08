package main

import "fmt"

func main() {
    a, b, c := 10, 15, 20
    fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

    if a > b {
        fmt.Println("a is greater than b")
    } else {
        fmt.Println("a is less than b")
    }

    if a > b && a > c {
        fmt.Println("a is bigger than b and c")
    } else if b > a && b > c {
        fmt.Println("b is bigger than a and c")
    } else if c > a && c > b {
        fmt.Println("c is bigger than a and b")
    }

    option := "A"

    // In Go, "break" is implicit on the defined cases, so you dont need to write it
    switch {
    case option == "A":
        fmt.Println("Option A selected!")
    case option == "B":
        fmt.Println("Option B selected!")
    case option == "C":
        fmt.Println("Option C selected!")
    default:
        fmt.Println("Unknown option provided!")
    }
}