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
}