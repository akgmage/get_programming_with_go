// The big.Int type can happily store and operate on the distance to Andromeda Galaxy, a mere 24 quintillion kilometers.
// Opting to use big.Int requires that you use it for everything in your equation, even the constants you had before.
// The NewInt function takes an int64 and returns a big.Int
// Div method performs the necessary division so the result can be displayed
package main

import (
	"fmt"
	"math/big"
)

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away") // Prints Andromeda Galaxy is 24000000000000000000 km away

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed") // Prints That is 926568346 days of travel at light speed
}
