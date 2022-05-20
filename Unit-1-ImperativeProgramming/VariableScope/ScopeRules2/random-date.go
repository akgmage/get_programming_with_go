// To remove duplication and simplify the code, the variables in listing(ScopeRules/scope-rules.go)
// should be declared in wider function scope, making them available after the switch statement for later work.
// Its time to refactor!
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
