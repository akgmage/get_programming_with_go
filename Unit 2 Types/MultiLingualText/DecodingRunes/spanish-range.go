// The Go language provides the range keyword to iterate over a variety of collections.
// It can also decode UTF-8 encoded strings
package main

import "fmt"

func main() {
	question := "¿cómo estás?"

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
	// If you dont need the index, the blank (an undersocre) allows you to ignore it
	for _, c := range question {
		fmt.Printf("%c", c) // Prints ¿cómo estás?
	}
	// Prints
	/*
		0 ¿
		2 c
		3 ó
		5 m
		6 o
		7
		8 e
		9 s
		10 t
		11 á
		13 s
		14 ?
	*/
}
