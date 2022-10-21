// Go’s select lets you wait on multiple channel operations. Combining goroutines
// and channels with select is a powerful feature of Go.
package main

import (
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
}