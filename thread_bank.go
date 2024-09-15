package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	alice := 10000
	bob := 10000
	var mu sync.Mutex

	total := alice + bob

	// incorrect code
	// take from alice and send it to bob
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice = alice - 1
			mu.Unlock()
			mu.Lock()
			bob = bob + 1
			mu.Unlock()
		}
	}()

	// take form bob send to alice
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob = bob - 1
			mu.Unlock()
			mu.Lock()
			alice = alice + 1
			mu.Unlock()
		}
	}()

	// correct code
	/*
		go func() {
			for i := 0; i < 1000; i++ {
				mu.Lock()
				alice = alice - 1
				bob = bob + 1
				mu.Unlock()
			}
		}()

		go func() {
			for i := 0; i < 1000; i++ {
				mu.Lock()
				bob = bob - 1
				alice = alice + 1
				mu.Unlock()
			}
		}()
	*/

	// below code detects violation, for correct code alice+bob should be == total
	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("violation alice=%d, bob=%d, total=%d\n", alice, bob, alice+bob)
		}
		mu.Unlock()
	}
}
