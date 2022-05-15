// Another way to arrive at true or false is by comparing two values.

package main

import "fmt"

func main() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 27
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}
