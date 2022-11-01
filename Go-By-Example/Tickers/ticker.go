package main

import (
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for  {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmr.Println("Tick at", t)
			}
		}
	}()
}