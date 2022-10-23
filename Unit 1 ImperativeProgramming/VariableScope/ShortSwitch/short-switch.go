// Short declaration can also be used as part of a switch statement

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("random spacetime #", num)
	}
}
