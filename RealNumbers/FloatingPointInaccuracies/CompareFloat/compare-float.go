// INstead of comparing floating-point numbers directly,
// determine the absolute difference between two numbers
// and then ensure the difference isn't too big
package main

import (
	"fmt"
	"math"
)

func main() {
	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank == 0.3) // Prints flase

	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) // Prints true
}
