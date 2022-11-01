// Implement a worker pool using goroutines and channels.
package main

import (
	"fmt"
	"time"
)
// Here’s the worker, of which we’ll run several concurrent 
// instances. These workers will receive work on the jobs channel and send the corresponding results on results. 
// We’ll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, result chan<- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finishedjob", j)
		result <- j * 2
	}
}