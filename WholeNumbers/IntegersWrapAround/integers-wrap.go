// Integers have limited range. When that range is exceeded, integer types in Go, wrap around
package main

import "fmt"

func main() {

	var red uint8 = 255
	red++
	fmt.Println(red) // Prints 0

	red = 255
	red += 2
	fmt.Println(red)

	red = 0
	red--
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number) // Prints -128

	number = -128
	number--
	fmt.Println(number)

	var green uint16 = 65535
	green++
	fmt.Println(green)

}
