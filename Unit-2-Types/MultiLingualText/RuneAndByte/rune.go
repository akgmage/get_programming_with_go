// To represent a single Unicode code point, Go provides rune, which is an alias for the int32 type.
package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) // Prints 960 940 969 33
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) // Prints π ά ω !
}
