// Literal strings and raw strings both result in strings
// Detailed description Source (https://pkg.go.dev/fmt)
package main

import "fmt"

func main() {
	// %T a Go-syntax representation of the type of the value
	// [SOME_NUMBER] refers to value in SOME_NUMBER position 
	fmt.Printf("%v is a %[1]T\n", "literal string", 2) // Prints literal string is a string
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)    // Prints raw string literal is a string
}
