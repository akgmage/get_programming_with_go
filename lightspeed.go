// Constatnts and variables can help by providing descriptive names
// How long does it take to get to mars?
package main

import "fmt"

func main() {
	const lightspeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightspeed, "seconds") // Prints 186 seconds

	distance = 401000000

	fmt.Println(distance/lightspeed, "seconds") // Prints 1337 seconds
}
