package main

import "fmt"

func foo() string {
	return "bar"
}

func main() {
	// Variable "name" with type "string"
	var name string = "Lucas"
	// note: "var" keyword is required when a data type is defined
	fmt.Printf("Hey, %s!\n", name)

	// Go doesn't need to explicitly define the type of an variable, you can just use the operator
	// ":=" and the compiler will infer the data type from the value that you have provided
	age := 16
	fmt.Printf("You have %dy.o., right?\n", age)

	// For define multiple variables at once:
	otherName, otherAge := "John Doe", 67
	fmt.Println("Other person:", otherName, otherAge, "\by.o.")

	// When calling a function inside a variable, its a good practice to define the expected data type
	var bar string = foo()
	fmt.Println("foo:", bar)
}