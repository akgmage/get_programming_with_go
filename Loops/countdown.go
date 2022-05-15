// Rather than type the same code multiple times, the for keyword repeats code for you.
// Before each iteration, the expression count > 0 is evaluated to produce a Boolean value.
// If value is false (count = 0), the loop terminates-- otherwise, it runs the body of the
// loop (the code between { and }).
package main

import (
	"fmt"
	"time"
)

func main() {
	var count = 10 // Declares and initializes

	for count > 0 { // A condition
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- // Decrements count; otherwise it will loop forever
	}
	fmt.Println("Liftoff!")
}
