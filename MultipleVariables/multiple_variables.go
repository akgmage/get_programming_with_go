// Declare multiple variables at once
package main

import "fmt"

func main() {

	var distance1 = 56000000
	var speed1 = 100800
	fmt.Println(distance1)
	fmt.Println(speed1)

	// Or you can declare them as group

	var (
		distance2 = 56000000
		speed2    = 100800
	)

	fmt.Println(distance2)
	fmt.Println(speed2)

}
