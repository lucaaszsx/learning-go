package main

import (
    "fmt"
    "strconv"
)

func main() {
    // ----- String to Number conversions

    // Atoi: convert string to int (most common case)
    num, err := strconv.Atoi("42")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Atoi '42':", num) // 42

    // Atoi with invalid input
    _, err = strconv.Atoi("abc")
    if err != nil {
        fmt.Println("Atoi 'abc' error:", err)
    }

    // ParseFloat: convert string to float64
    // second argument is the bit size (32 or 64)
    pi, err := strconv.ParseFloat("3.14159", 64)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("ParseFloat '3.14159':", pi)

    // ParseInt: convert string to integer with base and bit size
    // base 10 = decimal, base 16 = hexadecimal, base 2 = binary
    hex, _ := strconv.ParseInt("FF", 16, 64)
    fmt.Println("ParseInt 'FF' (hex):", hex) // 255

    binary, _ := strconv.ParseInt("1010", 2, 64)
    fmt.Println("ParseInt '1010' (binary):", binary) // 10

    // ParseBool: convert string to bool
    // accepts: "1", "t", "T", "TRUE", "true", "True" → true
    // accepts: "0", "f", "F", "FALSE", "false", "False" → false
    flag, _ := strconv.ParseBool("true")
    fmt.Println("ParseBool 'true':", flag)

    // ----- Number to String conversions

    // Itoa: convert int to string (most common case)
    str := strconv.Itoa(42)
    fmt.Println("Itoa 42:", str)          // "42"
    fmt.Printf("Type: %T\n", str)         // string

    // FormatFloat: convert float64 to string
    // 'f' = decimal format, -1 = shortest precision, 64 = bit size
    floatStr := strconv.FormatFloat(3.14159, 'f', 2, 64)
    fmt.Println("FormatFloat 3.14159 (2 decimals):", floatStr) // "3.14"

    // FormatInt: convert int64 to string with different bases
    fmt.Println("FormatInt 255 (hex):", strconv.FormatInt(255, 16))   // "ff"
    fmt.Println("FormatInt 10 (binary):", strconv.FormatInt(10, 2))   // "1010"

    // FormatBool: convert bool to string
    fmt.Println("FormatBool true:", strconv.FormatBool(true)) // "true"

    // ----- Quick conversions with Itoa/Atoi
    // These are the most used in everyday Go code
    input := "999"
    converted, _ := strconv.Atoi(input)
    converted++
    result := strconv.Itoa(converted)
    fmt.Println("999 + 1 =", result) // "1000"
}
