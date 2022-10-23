// You can use switch statement with conditions for each case, much like using if..else.
// One unique feature of switch is the {fallthrough} keyword, which is used to execute the body of the next case
package main

import "fmt"

func main() { // Expressions are in each case
	fmt.Println("There is a cavern entrance here and a patj to the east.")

	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems old enough.")
		fallthrough // Falls through to the next case
	case room == "underwater":
		fmt.Println("The water is freezing cold")
	}
}
