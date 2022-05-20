// IN mathematics , some rational numbers can't be accurately represented in decimal form.
// The number 0.33 is only an approximation of 1/3.
// Unsurprisingly, a calculation on approximate values has an approximate result.
// 1/3 + 1/3 + 1/3 = 1
// 0.33 + 0.33 + 0.33 = 0.99
// Floating-point numbers suffer from rounding errors too, except floating-point uses a binary representation (using only 0s and 1s)
// instead of using decimal(1-9). The consequence is that computers can accurately represent 1/3 but have rounding errors
// with other numbers.
package main

import "fmt"

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank)
}
