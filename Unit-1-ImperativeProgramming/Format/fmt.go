// The Prinr and Println functions have a sibling that gives more control over output.
// By using Printf, we can insert values anywhere in the text
package main

import "fmt"

func main() {
	// The first argument to Printf is always a text. The text contains
	// the format verb %v which is substituted for value of the expression provided by the second argument
	fmt.Printf("My weight on the surface of Mars is %v lbs", 187.3*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 27*365/687)

	// If multiple format verbs are specified, the Printf function will substitute multiple values in order.
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 187.3)

	// Printf can help you align text. Specify a width as part of the format verb, such as %4v to pad a value
	// to a width of 4 characters. A positive number pads with spaces to the left, and a negative number pads
	// with spaces to the right.
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	// The preceding code displays the following output:
	// SpaceX          $  94
	// Virgin Galactic $ 100
}
