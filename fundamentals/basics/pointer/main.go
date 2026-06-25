package main

import "fmt"

func main() {
    // A pointer stores the memory address of a variable, not the value itself
    // Use the "&" operator to get the address of a variable
    // Use the "*" operator to dereference (access the value at the address)
    x := 42
    p := &x // p is a pointer to x

    fmt.Println("Value of x:", x)       // 42
    fmt.Println("Address of x:", p)     // memory address (e.g., 0xc0000b2008)
    fmt.Println("Value via pointer:", *p) // 42

    // Modifying the value through the pointer
    *p = 100
    fmt.Println("New value of x:", x) // 100

    // Pointers are useful to avoid copying large structs
    // and to allow functions to modify the original variable
    age := 25
    increment(&age)
    fmt.Println("Age after increment:", age) // 26

    // Pointer to a pointer (double pointer)
    pp := &p
    fmt.Println("Value via double pointer:", **pp) // 100

    // The zero value of a pointer is nil
    var ptr *int
    fmt.Println("Nil pointer:", ptr) // <nil>

    // Always check for nil before dereferencing
    if ptr != nil {
        fmt.Println("Value:", *ptr)
    } else {
        fmt.Println("Pointer is nil, cannot dereference")
    }
}

// increment receives a pointer to an int and modifies the original value
func increment(n *int) {
    *n++
}
