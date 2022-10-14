package main

import "fmt"

/// Calls itself until it reaches base case of fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
		
}
func main() {
	fmt.Println(fact(5))
	// Closures can be recursive, this requires the closure to be 
	// declared with a typed var explicitly before it’s defined
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n - 2)
	}
	fmt.Println(fib(10))
}