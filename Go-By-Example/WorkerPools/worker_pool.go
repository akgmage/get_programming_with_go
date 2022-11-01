package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
	}
}