// 64-bit vs 32-bit floating-point
package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64) // Prints 3.141592653589793
	fmt.Println(pi32) // Prints 3.1415927
}
