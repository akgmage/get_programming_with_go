// A function that returns a Boolean value
// The following listing uses  the --Contains-- function froms strings package
// to ask if the {command} variable contains the text "outside"
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit) // Print You leave the cave: true

}
