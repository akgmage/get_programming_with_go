// If you're ever curious about which type the Go compiler inferred, the Printf function provides
// the %T verb to display a variable's type
package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year) // Prints type int for 2018

	// Instead of repeating the variable twice, you can tell Printf to use the first argument [1]
	// for the second format verb

	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days) // Prints type int for 365.2425

	a := "text"
	fmt.Printf("Type %T for %[1]v", a)
}
