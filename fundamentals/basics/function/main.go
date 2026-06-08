package main

import (
    "errors"
    "fmt"
)

func main() {
    var a, b int = 10, 20
    var sum = sumTwo(a, b)
    fmt.Println("Result of sumTwo:", sum)

    var max int = 25
    var safeSum, err = safeSumTwo(a, b, max)

    if err != nil {
        fmt.Println(err.Error())
        // or: panic(err.Error())
        return
    }
    
    fmt.Println("safeSumTwo doesn't returned an error! Result:", safeSum)
}

// Simple function with 2 integer arguments and a single return value 
func sumTwo(a int, b int) int {
    return a + b
}

// safeSumTwo receives a and b as int and returns an integer value *and* an error
func safeSumTwo(a int, b int, max int) (int, error) {
    var err error
    sum := a + b

    if sum > max {
        // fmt.Sprintf: useful to format a message with various replacement values
        msg := fmt.Sprintf("The sum between %d and %d is greater than the maximum value of %d", a, b, max)
        // errors.New: creates an error with a custom message
        err = errors.New(msg)
        return 0, err
    }

    return sum, err
}