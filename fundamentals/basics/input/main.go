package main

import "fmt"

func main() {
	// For getting user input, use `fmt.Scan`, `fmt.Scanf`, `fmt.Scanln`, etc.
	var x, a, b int
	var name string

	fmt.Print("What's your name?: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! How are you?\n", name)

	fmt.Print("Enter the value of X: ")
	fmt.Scan(&x)
	fmt.Println("You entered:", x)

	fmt.Print("Enter two numbers separated by colon: ")
	fmt.Scanf("%d, %d", &a, &b)
	fmt.Printf("Sum between %d and %d = %d\n", a, b, a + b)
}

// see: bufio standard Go library