package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// video link : https://www.youtube.com/watch?v=NIvSQCwcots

var ready bool = false

func main() {
	// gettingReadyForMission()
	gettingReadyForMissionWithCond() // reimplementing above function with shared condition
}

func gettingReadyForMission() {
	go gettingReady()
	workIntervals := 0

	for !ready {
		workIntervals++
	}
	fmt.Printf("We are now ready! After %d workIntervals\n", workIntervals)
}

func gettingReadyForMissionWithCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond) // not &cond as above line NewCond already returns cond address or &cond
	workIntervals := 0

	cond.L.Lock()
	for !ready {
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("We are now ready! After %d workIntervals\n", workIntervals)
}

func gettingReady() {
	sleep() // sleep randomly between 1 and 5 seconds
	ready = true
}

func gettingReadyWithCond(cond *sync.Cond) {
	sleep() // sleep randomly between 1 and 5 seconds
	ready = true
	cond.Signal()
	// Signal because single coroutine, if mutiple then we use Broadcast the same way
	// Signal wakes up only 1 thread that is waiting while Broadcast wakes up all thread that are waiting
}

// function to sleep for random seconds between 1 and 5
func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
