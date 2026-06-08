package main

import "fmt"

const myConstant string = "some_cool_value"
// or just: const myConstant = ...

func main() {
    // myConstant = "some_other_value" -> compiler returns an error: cannot assign to myConstant...
    fmt.Println(myConstant)
}