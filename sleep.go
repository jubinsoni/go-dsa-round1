package main

import (
	"fmt"
	"sync"
	"time"
)

// Main Lesson : Don't use your brain, if varaible(here "done") is shared between
// Coroutine(here between main and periodic coroutine) then use Locks whenever we are doing read/write operation

var done bool
var mu sync.Mutex

func main() {
	time.Sleep(1 * time.Second)
	fmt.Println("started")
	go Periodic()
	time.Sleep(5 * time.Second) // wait for a while so that we can observe what ticker does
	mu.Lock()
	done = true
	mu.Unlock()

	fmt.Println("cancelled")
	time.Sleep(3 * time.Second)
}

func Periodic() {
	for {
		fmt.Println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			return
		}
		mu.Unlock()
	}
}
