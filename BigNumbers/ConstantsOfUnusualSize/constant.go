// Under the hood, untyped numeric constants are backed by the big package,
// enabling all the usual operations with numbers well beyond 18 quintillion
// For variables, Go uses type inference to determine the type, and in the case of 24 quintillion, overflows int type.
// Constants are different. rather than infer a type, constants can be untyped.
// const distance = 24000000000000000000 doesn't cause an overflow error
// Constant are declared with const keyword, but every literal value in your program is a constant too.
// This means unusually sized numbers can be used directly.
package main

import "fmt"

func main() {

	fmt.Println("Andromeda Galaxy is", 24000000000000000000/299792/86400, "light days away") // // Prints Andromeda Galaxy is 926568346 light days away.

	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400

	const days = distance / lightSpeed / secondsPerDay

	fmt.Println("Andromeda Galaxy is", days, "light days away.") // Prints Andromeda Galaxy is 926568346 light days away.
}
