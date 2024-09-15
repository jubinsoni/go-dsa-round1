package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func main() {
	var num_producers int = 4

	c := make(chan int)

	for i := 0; i < num_producers; i++ {
		go producer(c, i)
	}
	// go producer(c)
	go consumer(c)

	// run producer-consumer code for 10 seconds
	time.Sleep(10 * time.Second)
}

func producer(c chan int, i int) {
	for {
		time.Sleep(1 * time.Second)
		// c <- rand.Int()
		c <- i
	}
}

func consumer(c chan int) {
	for {
		v := <-c
		fmt.Println(v)
	}
}
