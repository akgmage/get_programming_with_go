// range iterates over elements in a variety of data structures
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// range on arrays and slices provides both the index and value for each entry
	// ignore index with blank identifier _
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	
	kvs := map[string]string{"a": "apple", "b":"banana"}
	// range on map iterates over key/value pairs.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}