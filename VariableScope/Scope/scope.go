// In go scope tend to begin and end along the lines of curly braces {}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 { // A new scope begins
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // Scope ends
}
