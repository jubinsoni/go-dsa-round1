package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MIT video link : https://youtu.be/UzzcUS2OHqo?si=GycGGzkq1vpzBHGA&t=1414
// we are using Condition variables here

// At an high levl pattern looks like this
/*
thread A
mu.Lock()
// do something to change condition
cond.Broadcast()
mu.Unlock()

-------------------
thread B(or main thread)

mu.Lock()
for {
  if condition
  cond.wait()
}
//now condition is true and we have the lock
mu.Unlock()
*/

func main() {
	//basicRaftImpl()
	basicRaftImplWithBroadCast()
}

func basicRaftImplWithBroadCast() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu) // condition variables are associated with Locks, so it is given address to the lock during initialization

	for i := 0; i < 10; i++ {
		go func() {
			// vote := requestVote() // simulating master machine asks for vote from 10 slave machine
			// as of now we bypass to true i.e each slave machine gave a vote
			var vote bool = true
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count = count + 1
			}
			finished++
			cond.Broadcast()
		}()
	}

	mu.Lock()
	for {
		if count >= 5 || finished == 10 {
			break
		}
		cond.Wait()
	}
	mu.Unlock()
	if count >= 5 {
		fmt.Println("received 5+ votes")
	} else {
		fmt.Println("lost")
	}
}

func basicRaftImpl() {
	rand.Seed(time.Now().UnixNano())

	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			// vote := requestVote() // simulating master machine asks for vote from 10 slave machine
			// as of now we bypass to true i.e each slave machine gave a vote
			var vote bool = true
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count = count + 1
			}
			finished++
		}()
	}

	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()
		// we are sleeping here for 50 Millisecond so that
		// a machine core don't get cclogged by infinite for loop
		time.Sleep(50 * time.Millisecond)
	}

	if count >= 5 {
		fmt.Println("received 5+ votes")
	} else {
		fmt.Println("lost")
	}
	// mu.Unlock()
}
