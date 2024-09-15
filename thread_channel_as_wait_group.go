package main

import (
	"fmt"
)

func main() {
	var num_routines int = 5
	done_ch := make(chan bool)
	for i := 0; i < num_routines; i++ {
		go func(x int) {
			sendRPC(x)
			done_ch <- true
		}(i)
	}

	for i := 0; i < num_routines; i++ {
		<-done_ch
	}
}

func sendRPC(i int) {
	fmt.Println(i)
}
