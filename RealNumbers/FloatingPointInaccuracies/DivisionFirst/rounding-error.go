// To minimize rounding errors, we recommend that you perform multiplication before division.
// The result tends to be more accurate that way.
package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, " F\n") // Prints 69.80000000000001 F
	fmt.Print((9.0/5.0*celsius)+32, " F\n") // Prints 69.80000000000001 F
}
