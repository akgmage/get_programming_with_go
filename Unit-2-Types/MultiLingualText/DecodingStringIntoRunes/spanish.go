// The first step to supporting other languages is to decode characters to the rune type before manipulating them.
// Fortunately, Go has functions and language features for decoding UTF-8 encoded strings
// The utf package provides functions to determine the length of string in runes rather than bytes
// and to decode the first character of a string
// The DecodeRuneInString function returns the first character and the number of bytes the character consumed.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿cómo estás?"

	fmt.Println(len(question), "bytes")                    // Prints 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // Prints 12 runes

	// NOTE:  Unlike many programming languages, functions in Go can return multiple values.

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) // Prints First rune: ? 1 bytes

}
