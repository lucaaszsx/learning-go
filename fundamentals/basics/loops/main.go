package main

import "fmt"

// Go only have "for" to use in loop structures
func main() {
	// --- Simple
	var count int

	fmt.Print("Enter a number: ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		fmt.Println("Count:", i)
	}

	for i := count; i >= 1; i-- {
		fmt.Println("Reverse count:", i)
	}

	// --- List
	letters := []string{"a", "b", "c"}
	for index, v := range letters {
		fmt.Printf("%d° letter: %s\n", index, v)
	}

	// --- Map
	urls := map[string]string{
		"Lucas": "https://github.com/lucaaszsx",
		"Go": "https://github.com/golang",
	}

	// Order of "range" in maps aren't granted
	for name, url := range urls {
		fmt.Println(name, "=", url)
	}

	// --- While-style
	var opt string

	loop: for opt != "0" {
		fmt.Println("Select an option (or 0 to exit): A, B, C")
		fmt.Scan(&opt)

		switch opt {
		case "A":
			fmt.Println("Hello, world! (option A)")
		case "B":
			fmt.Println("Hi, user! (option B)")
		case "C":
			fmt.Println("Hello! (option C)")
		case "0":
			fmt.Println("Program ended. Breaking loop...")
			break loop
		default:
			fmt.Println("Invalid option. Must be one of: A, B, C, or 0 to exit.")
		}
	}

	// --- Infinite loop
	/*x := 0

	for {
		fmt.Println(x)
		x++
	}*/
}