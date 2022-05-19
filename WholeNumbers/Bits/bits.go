// The %b format verb will show you the bits for an integer value.
// Like other format verbs, %b can be zero padded to a minimum length
package main

import "fmt"

func main() {
	var green uint8 = 3
	fmt.Printf("%06b\n", green) // Prints 00000011

	green++
	fmt.Printf("%06b\n", green) // Prints 00000100
}
