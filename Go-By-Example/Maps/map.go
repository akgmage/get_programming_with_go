package main

import "fmt"

func main() {
	// create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)
	// Set key/value pairs using typical name[key] = val syntax
	m["item1"] = 7
	m["item2"] = 9

	fmt.Println("map:", m)
	// len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))
	x := m["item1"]
	fmt.Println("x:",x)
	delete(m, "item2")
	fmt.Println("map", m)

	_, present := m["item2"]
	fmt.Println("present:",present)
}