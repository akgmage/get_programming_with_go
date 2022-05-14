package main

import "fmt"

func main() {
	// Line 7 and 8 are equivalent
	var weight = 187.3
	weight = weight * 0.3783
	weight *= 0.3783
	fmt.Println(weight)
	// Incrementing by 1 has an additional shoutcut

	var age = 27
	age = age + 1
	age += 1
	age++
	fmt.Println(age)
	// Go doen not support the prefix increment ++count like in C and Java
}
