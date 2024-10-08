package main

import (
	"fmt"
	"time"
)

// Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go
// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.
// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
// We receive the values "one" and then "two" as expected.
// Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

