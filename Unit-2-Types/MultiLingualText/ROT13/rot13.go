// The following listing will decipher a message from space
package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fngvba"

	for i := 0; i < len(message); i++ { // Iterates through each ASCII character
		c := message[i]
		if c >= 'a' && c <= 'z' { // Leaves spaces and punctuation as they are
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c) // Prints hi international space sation
	}
}
