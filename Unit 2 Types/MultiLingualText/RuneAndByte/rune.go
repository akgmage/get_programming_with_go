// To represent a single Unicode code point, Go provides rune, which is an alias for the int32 type.
package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	// %v the value in a default format
	// %c the character represented by the corresponding Unicode code point
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang) // Prints 960 940 969 33
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang) // Prints π ά ω !

	// Rather than memorize Unicode code points, Go provides a character literal,
	// Just enclose a character in a single quotes 'A'. If no type is specified.
	// Go will infer a rune. So following three lines are equivalent.
	// grade := 'A'
	// var grade = 'A'
	// var grade rune = 'A'
	// The grade value still contains a numeric value, in this case 65,
	// the code point for capital letter 'A'
	// Character literals can also be used with the byte alias:
	// var byte = "*"
}
