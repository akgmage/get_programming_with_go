// Decipher code from Julius Caesar

package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg abc"

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'z' {
			c = c - 3 // shift 3 letters backwards
			if c < 'a' && c > 'Z' {
				c = c + 26
			}
		}
		fmt.Printf("%c", c) // Prints I came, I saw, I conquered
	}
}
