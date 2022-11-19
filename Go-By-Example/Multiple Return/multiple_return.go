/// Go has built-in support for multiple return values. This feature is used often in idiomatic Go,
/// for example to return both result and error values from a function
package main

import "fmt"

// (int, int) in this function signature shows that the function returns 2 ints
func vals(a int, b int) (int, int) {
	return a, b
}
func main() {
	// Here we use the 2 different return values from the call with multiple assignment
	a, b := vals(1, 2)

	fmt.Println(a)
	fmt.Println(b)
	// use the blank identifier (_) if you only want a subset of the returned values, 
	_, c := vals(2, 3)
	fmt.Println(c)
}