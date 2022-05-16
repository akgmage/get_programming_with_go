// The followint listing demonstrates a variant of the for loop that combines initialization, a condition, and a post statement that decrements count.
package main

import "fmt"

func main() {
	var count = 0
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count) // count remains in scope
}
