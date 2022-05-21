// Your program can access individual characters, but it can't alter the characters of a string.
// The following listing uses square brackets [] to specify an index into a string which accesses
// a single byte (ASCII character). The index starts from zero

package main

import "fmt"

func main() {
	peace := "shalom"
	peace = "salam"
	fmt.Println(peace) // Prints salam

	message := "shalom"
	c := message[5]
	fmt.Printf("%c\n", c) // Prints m
}
