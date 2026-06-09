package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b float64
	fmt.Print("Enter the numerator and denominator separated by a comma: ")
	fmt.Scanf("%f, %f", &a, &b)

    result, err := divide(a, b)
	if err != nil {
		fmt.Println("divide returned an error:", err)
		return
	}

	fmt.Printf("%f / %f = %f\n", a, b, result)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	
	return a / b, nil
}