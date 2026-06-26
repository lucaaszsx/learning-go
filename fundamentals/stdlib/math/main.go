package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    // ----- Constants from math package
    fmt.Println("Pi:", math.Pi)       // 3.141592653589793
    fmt.Println("E:", math.E)         // 2.718281828459045
    fmt.Println("Phi:", math.Phi)     // 1.618033988749895
    fmt.Println("Sqrt2:", math.Sqrt2) // 1.4142135623730951

    // ----- Basic math operations
    fmt.Println("Abs(-10):", math.Abs(-10))         // 10
    fmt.Println("Max(3, 7):", math.Max(3, 7))       // 7
    fmt.Println("Min(3, 7):", math.Min(3, 7))       // 3

    // ----- Powers and roots
    fmt.Println("Pow(2, 10):", math.Pow(2, 10))     // 1024
    fmt.Println("Sqrt(144):", math.Sqrt(144))        // 12
    fmt.Println("Cbrt(27):", math.Cbrt(27))          // 3
    fmt.Println("Log(100):", math.Log(100))          // natural log (ln)
    fmt.Println("Log10(100):", math.Log10(100))      // log base 10 → 2
    fmt.Println("Log2(8):", math.Log2(8))            // log base 2 → 3

    // ----- Rounding
    fmt.Println("Floor(3.7):", math.Floor(3.7))     // 3 (round down)
    fmt.Println("Ceil(3.2):", math.Ceil(3.2))       // 4 (round up)
    fmt.Println("Round(3.5):", math.Round(3.5))     // 4 (round to nearest)
    fmt.Println("Round(3.4):", math.Round(3.4))     // 3

    // ----- Trigonometry (angles in radians)
    fmt.Println("Sin(π/2):", math.Sin(math.Pi/2))   // 1
    fmt.Println("Cos(0):", math.Cos(0))             // 1
    fmt.Println("Tan(π/4):", math.Tan(math.Pi/4))   // ~1

    // ----- Special values
    fmt.Println("IsNaN(math.NaN()):", math.IsNaN(math.NaN()))     // true
    fmt.Println("IsInf(+Inf, 1):", math.IsInf(math.Inf(1), 1))     // true (positive infinity)

    // ----- Random numbers
    // rand.Intn returns a random int in [0, n)
    fmt.Println("Random 1-100:", rand.Intn(100)+1)

    // rand.Float64 returns a random float in [0.0, 1.0)
    fmt.Println("Random float:", rand.Float64())

    // rand.Intn with a seed for reproducible results (useful for testing)
    // Note: in Go 1.20+, the global rand functions are automatically seeded

    // ----- Practical example: distance between two points
    x1, y1 := 0.0, 0.0
    x2, y2 := 3.0, 4.0
    distance := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
    fmt.Printf("Distance from (%.0f,%.0f) to (%.0f,%.0f): %.2f\n", x1, y1, x2, y2, distance) // 5.00
}
