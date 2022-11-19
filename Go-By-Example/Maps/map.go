package main

import "fmt"

func main() {
	// create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)
	// Set key/value pairs using typical name[key] = val syntax
	m["item1"] = 7
	m["item2"] = 9

	mp := make(map[int]int)
	mp[1] = 2
	mp[2] = 1
	fmt.Println(mp)

	fmt.Println("map:", m)
	// len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))
	x := m["item1"]
	fmt.Println("x:",x)
	delete(m, "item2")
	fmt.Println("map", m)

	// optional second return value when getting a value from a map indicates if the key was present in the map
	_, present := m["item2"]
	fmt.Println("present:",present)


	// declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	p := map[string]string{"foo": "bar", "bar": "foo"}
	fmt.Println(p)
}