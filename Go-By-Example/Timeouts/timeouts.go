// Timeouts are important for programs that connect to
// external resources or that otherwise need to bound execution time.
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	// suppose we’re executing an external call that returns its result on a channel c1 after 2s.
	// Note that the channel is buffered, so the send in the goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")		
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")	
	}
}