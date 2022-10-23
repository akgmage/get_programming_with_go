// My weight loss program. (A comment for human readers)
// Program to calculate how young and light I would be on Mars
/*
	Mars takes 687 Earth days to travel around the sun,
	and its weaker gravitational force means everything
	weighs approximately 38% of what it does on earth.
*/
package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(187.3 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(27 * 365 / 687)
	fmt.Print(" years old.")
}
