package main

import (
	"fmt"
	"sync"
	"time"
)

// simple rule: Don't use brain, whenever a variable which is used inside goroutine
// used anywhere else(including the main coroutine) then always use mutex to read/write that variable

func main() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter = counter + 1
		}()
	}

	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Println(counter)
	mu.Unlock()
}
