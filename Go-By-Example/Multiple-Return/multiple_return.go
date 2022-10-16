/// Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
/// for example to return both result and error values from a function
package main

import "fmt"

// (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 1, 2
}
func main() {
	// Here we use the 2 different return values from the call with multiple assignment
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)
	// use the blank identifier (_) if you only want a subset of the returned values, 
	_, c := vals()
	fmt.Println(c)
}