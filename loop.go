package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(x int) {
			SendRPC(x)
			wg.Done()
		}(i) // passing i as argument for goroutine inline function

		// below is wrong way of doing if goroutine directly uses variable i directly from outsied instead
		// see video for more clarity https://youtu.be/UzzcUS2OHqo?si=Fy4Xa5Fmx2PXUicX&t=293
		// go func() {
		// 	SendRPC(i)
		// 	wg.Done()
		// }()
	}
	wg.Wait()
}

func SendRPC(x int) {
	fmt.Println(x)
}
