// A goroutine is a lightweight thread of execution
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}