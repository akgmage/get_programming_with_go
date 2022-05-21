package main

import "fmt"

func main() {
	fmt.Printf("%v is a %[1]T\n", "literal string", "asd")
	fmt.Printf("%v is a %[1]T\n", "raw string literal")
}
