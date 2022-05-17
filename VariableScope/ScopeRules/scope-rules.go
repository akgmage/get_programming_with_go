// The code in the following listing generates and displays a random date -- perhaps a departure
// date to mars. It also demonstrates several different scopes in Go and shows why considering
// scope when declaring variables is important

package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // era variable is available throughout the package

func main() {
	year := 2022                               // era and year are in scope
	switch month := rand.Intn(12) + 1; month { // era, year, month are in scope
	case 2:
		day := rand.Intn(28) + 1 // era, year, month, day are in scope
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // Its a new day
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // Its a new day
		fmt.Println(era, year, month, day)
	} // month and day are out of scope
} // year is no longer in scope
