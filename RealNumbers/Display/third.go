// The %f verb formats the value of third with a width and with precision [%4.2f] (4 is the width, and 2f means Precision,
// how many digits (2) should appear after the decimal point)
// Width specifies minimum numbers of characters to display
package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)           // Prints 0.3333333333333333
	fmt.Printf("%v\n", third)    // Prints 0.3333333333333333
	fmt.Printf("%f\n", third)    // Prints 0.333333
	fmt.Printf("%.3f\n", third)  // Prints 0.333
	fmt.Printf("%4.2f\n", third) // 0.33
}
