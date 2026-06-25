package main

import (
    "fmt"
    "math"
)

// Shape defines a set of methods that a type must implement
// Any type with Area() and Perimeter() methods satisfies this interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle implements the Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Rectangle also implements the Shape interface
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// printShapeInfo accepts any type that implements the Shape interface
func printShapeInfo(s Shape) {
    fmt.Printf("Type: %T\n", s)
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
    fmt.Println("---")
}

// Stringer is a common interface in Go's fmt package
// When implemented, fmt.Println uses this method to format the output
func (c Circle) String() string {
    return fmt.Sprintf("Circle(%.2f)", c.Radius)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}

    // Both types satisfy the Shape interface
    printShapeInfo(circle)
    printShapeInfo(rectangle)

    // Using the Stringer interface (String() method)
    fmt.Println(circle)    // prints: Circle(5.00)
    fmt.Println(rectangle) // prints: Rectangle(4.00 x 6.00)

    // Interface as a collection of different types
    shapes := []Shape{circle, rectangle}
    for _, s := range shapes {
        fmt.Printf("Shape: %s, Area: %.2f\n", s, s.Area())
    }

    // Empty interface can hold any type (similar to any/Object in other languages)
    var anything interface{} = "hello"
    fmt.Println("Empty interface value:", anything)
    anything = 42
    fmt.Println("Empty interface value:", anything)

    // Type assertion to extract the underlying type
    value := anything.(int)
    fmt.Println("Asserted int value:", value)

    // Type assertion with check (safe way)
    if str, ok := anything.(string); ok {
        fmt.Println("It's a string:", str)
    } else {
        fmt.Println("It's not a string")
    }
}
