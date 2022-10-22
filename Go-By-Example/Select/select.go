// Go’s select lets you wait on multiple channel operations. Combining goroutines
// and channels with select is a powerful feature of Go.
package main

import (
	"fmt"
	"time"
)
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
// Each channel will receive a value after some amount of time, 
// to simulate 
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i:= 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)	
		}
	}
}