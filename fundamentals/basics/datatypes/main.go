package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	// ----- Integers: numbers without floating point
	// var myInt int = ... (int32 on 32-bit systems, int64 on 64-bit systems, etc.)
	// var myInt8 int8 = ... (-2^7 to 2^7-1)
	// var myInt16 int16 = ... (-2^15 to 2^15-1)
	// var myInt32 int32 = ... (-2^31 to 2^31-1)
	// var myInt64 int64 = ... (-2^63 to 2^63-1)
	year := 2026 // default to "int" type
	intA, intB := 10, 3
	result := intA / intB // inexact divisions with integers result in the outcome being rounded down

	fmt.Println("Current year:", year)
	fmt.Println("intA / intB =", result)

	// ----- Integers (unsigned): numbers without floating point and unsigned (always >=0)
	// var myUint uint = ... 
	// var myUint8 uint8 = ... (0 to 2^8-1)
	// var myUint16 uint16 = ... (0 to 2^16-1)
	// var myUint32 uint32 = ... (0 to 2^32-1)
	// var myUint64 uint64 = ... (0 to 2^64-1)
	var r, g, b uint8 = 255, 0, 232
	fmt.Printf("Color RGB: (%d, %d, %d)\n", r, g, b)

	// ----- Floats: floating-point numbers
	// var myFloat32 float32 = ... (IEEE 754, 32-bit, 6-9 digits of precision)
	// var myFloat64 float64 = ... (IEEE 754, 64-bit, 15-17 digits of precision, uses more memory)
	price := 34.9 // the inferred default type is "float64", so it's preferable to set it to "float32" whenever possible
	// note: it is not recommended to define monetary amounts using float as the data type
	fmt.Println("Price:", price)

	// ----- Strings (string): chain of characters
	myStr := "Hello, world!"
	lambda := "λ"

	fmt.Println("myStr:", myStr)
	// standard "len" func: when used in a string, returns the number of bytes, so it may not work as expected in some situations
	fmt.Println("Number of bytes in \"myStr\":", len(myStr)) // returns 13, its the real length of "myStr" in this case
	fmt.Printf("Number of bytes in \"%s\" (lambda): %d\n", lambda, len("λ")) // returns 2, cuz the lambda symbol has 2 bytes, but its just one character

	fmt.Println("Real length of \"myStr\":", utf8.RuneCountInString(myStr))
	fmt.Printf("Real length of \"%s\" (lambda): %d\n", lambda, utf8.RuneCountInString("λ"))

	// ----- Booleans: true or false, simply
	// var myBool bool = ...
	myBool := 1 > 2
	fmt.Println("myBool:", myBool)

	// ----- Default values of variables defined without a initial value depends on the data type
	// int, uint, float, rune: 0
	// string: '' (empty string)
	// bool: false
	var myDefInt int
	var myDefStr string
	var myDefBool bool

	fmt.Println("myDefInt:", myDefInt) // 0
	fmt.Println("myDefStr:", myDefStr) // ''
	fmt.Println("myDefBool:", myDefBool) // false
}