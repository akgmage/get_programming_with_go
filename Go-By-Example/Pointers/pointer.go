package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}
/// takes an int pointer [iptr]
func zeroptr(iptr *int) {
	// assigning a value to a dereferenced pointer changes the value 
	// at the referenced address.
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)
	zeroval(i)
	fmt.Println("zeroval", i)
	zeroptr(&i)
	fmt.Println("Pointer", i)
	// &i gives the memory address of i
	fmt.Println("Pointer address", &i)

}