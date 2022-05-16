// Short declaration makes it possible to declare a new variable in an if statement.
// IN the following listing the num variable can be used in any branch of the if statement
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
