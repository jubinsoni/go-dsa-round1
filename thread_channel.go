package main

import (
	"fmt"
	"time"
)

// video link : https://youtu.be/UzzcUS2OHqo?si=FFaeiu4MV3OSoqYN&t=2110

// channels are not  queue
// channels are syncronous in nature, basically thread safe
// send will only happen when someone is reading/receiving it, at this point synchronously they will exhange data over receiver
// if no one is reading/receiving data at other end then corresponding thread will get blocked which is trying to send data
// similarly
// read/receive will only happen when someone is sending it, at this point synchronously they will exhange data over receiver
// if no one is sending data at other end then corresponding thread will get blocked which is trying to read data

func main() {
	//writeToChannelBlock()
	// readFromChannelBlock()
	deadLockCode()
}

func writeToChannelBlock() {
	c := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		<-c // reads from channel
	}()

	start := time.Now()
	//main goroutine gets blocked until other goroutine receives
	c <- true // write to channel
	//
	fmt.Printf("send took %v\n", time.Since(start))
}

func readFromChannelBlock() {
	c := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		c <- true // write to channel
	}()

	start := time.Now()
	//main goroutine gets blocked until other goroutine writes
	<-c // read from channel
	//
	fmt.Printf("send took %v\n", time.Since(start))
}

func deadLockCode() {
	go func() {
		for {

		}
	}()
	c := make(chan bool)
	c <- true
	<-c
}
