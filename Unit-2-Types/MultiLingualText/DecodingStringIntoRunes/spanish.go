// The first step to supporting other languages is to decode characters to the rune type before manipulating them.
// Fortunately, Go has functions and language features for decoding UTF-8 encoded strings
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	question := "¿cómo estás?"

	fmt.Println(len(question), "bytes")                    // Prints 15 bytes
	fmt.Println(utf8.RuneCountInString(question), "runes") // Prints 12 runes

	c, size := utf8.DecodeLastRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size) // Prints First rune: ? 1 bytes

}
