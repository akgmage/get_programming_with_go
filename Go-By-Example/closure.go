package main

import "fmt"

/// This function intSeq returns another function, which we define 
/// anonymously in the body of intSeq. 
/// The returned function closes over the variable i to form a closure
func intseq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// This function value captures its own i value, which will be 
	// updated each time we call nextInt
	nextInt := intseq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intseq()
	fmt.Println(newInts())

}