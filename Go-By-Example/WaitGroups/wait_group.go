// To wait for multiple goroutines to finish, we can use a wait group.
package main

import (
	"fmt"
	"sync"
	"time"
)

// function we’ll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) //Sleep to simulate an expensive task
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines 
	// launched here to finish
	var wg sync.WaitGroup
	// Launch several goroutines and increment the WaitGroup counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i // Avoid re-use of the same i value in each goroutine closure
		// Wrap the worker call in a closure that makes sure to tell 
		// the WaitGroup that this worker is done. This way the worker itself 
		// does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()

}