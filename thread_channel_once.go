package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool = false

func main() {
	// initiateMission() // "marking mission completed" is called multiple times which is wrong
	initiateMissionOnce() // "marking mission completed" is called once which is right
}

func initiateMissionOnce() {
	var wg sync.WaitGroup
	wg.Add(100)

	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				once.Do(markMissionCompleted)
			}
			wg.Done()
		}()

	}
	wg.Wait()

	checkMissionCompleted()
}

func initiateMission() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				markMissionCompleted()
			}
			wg.Done()
		}()

	}
	wg.Wait()

	checkMissionCompleted()
}

func checkMissionCompleted() {
	if missionCompleted {
		fmt.Println("Mission is now completed")
	} else {
		fmt.Println("Mission was a failure")
	}
}

func markMissionCompleted() {
	missionCompleted = true
	fmt.Println("marking mission completed")
}

// we are getting a random number from 0-9 and approx. 1 goroutine would be able to find 0 == 0
func foundTreasure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
