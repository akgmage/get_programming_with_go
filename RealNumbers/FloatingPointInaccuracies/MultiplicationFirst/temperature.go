// To minimize rounding errors, we recommend that you perform multiplication before division.
// The result tends to be more accurate that way.
package main

import "fmt"

func main() {
	celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit, " F") // Prints 69.8 F
}

// What is the best way to avoid rounding errors ?
// Don't use floating point
