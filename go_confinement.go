package main

import (
	"fmt"
	"sync"
	"time"
)

// yt video followed :  https://www.youtube.com/watch?v=Bk1c30avsuU

/*
var mu sync.Mutex

func processDataMutex(wg *sync.WaitGroup, result *[]int, data int) {
	defer wg.Done()

	var temp int

	temp = data * 2
	mu.Lock()
	*result = append(*result, temp)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for {
		input := []int{1, 2, 3, 4, 5}
		result := []int{}
		// here 5 go routine are spawned and each concurrently tries to append
		// data to result, hence result array will have inconsistent length
    // to resolve this problem we use mutex so that only thread appends data at a time
    // thus making result consistent in length
		for _, data := range input {
			wg.Add(1)
			go processDataMutex(&wg, &result, data)
		}

		wg.Wait()
		time.Sleep(200 * time.Millisecond)

		fmt.Println(result)
	}
}
*/

func processDataConfinement(wg *sync.WaitGroup, result_address *int, data int) {
	defer wg.Done()

	var temp int

	temp = data * 2
	*result_address = temp
}

func main() {
	var wg sync.WaitGroup

	for {
		input := []int{1, 2, 3, 4, 5}
		result := make([]int, len(input))
		// basically we want to confine go routine to specific part of sahred resource
		// we should be able to write concurrently data to shared resource
		// so each individual go routine neeeds to be confined to an index in the slice
		// to do that send memory address and store data directly result to that address
		// this way each individual go routine access and modify a particular memory address
		for ind, data := range input {
			wg.Add(1)
			go processDataConfinement(&wg, &result[ind], data)
		}

		wg.Wait()
		time.Sleep(200 * time.Millisecond)

		fmt.Println(result)
	}
}
