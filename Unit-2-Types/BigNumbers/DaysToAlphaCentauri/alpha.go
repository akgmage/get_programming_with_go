// Rather than painstakingly typing every zero, you can write such numbers ig Go with exponent
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12
	fmt.Println("Alpha Centauri is", distance, "km away.") // Prints Alpha Centauri is 41300000000000 km away.

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed.") // Prints That is 1594 days of travel at light speed.

}
