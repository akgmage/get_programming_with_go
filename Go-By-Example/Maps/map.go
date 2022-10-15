package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["item1"] = 7
	m["item2"] = 9

	fmt.Println("map:", m)

	x := m["item1"]
	fmt.Println("x:",x)
	delete(m, "item2")
	fmt.Println("map", m)

	_, present := m["item2"]
	fmt.Println("present:",present)
}