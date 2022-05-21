// Literal strings and raw strings both result in strings
package main

import "fmt"

func main() {
	fmt.Printf("%v is a %[1]T\n", "literal string", "asd") // Prints literal string is a string
	fmt.Printf("%v is a %[1]T\n", `raw string literal`)    // Prints raw string literal is a string
}
