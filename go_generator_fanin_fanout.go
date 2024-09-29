package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// concepts covered: generics, pipeline, generator, fanin, fanout concurrency pattern
// yt video followed: https://youtu.be/wELNUHb3kuA?si=tq0FlnH74gEI8-BW&t=474

// generics is used because want any type of function to be passed here
// and then the channel gets instantiated based on the return type of function being called "make(chan T)"
// i.e if function returns int then an int channel will be instantiated
// if function return string then a string channel will be instantiated

// stages in pipeline used : genertor ===fanout-chan=== 3*findPrimes-stage ===3*fanin-chan=== fanin-stage ===chan=== taken-stage

func fanIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	// below transfer function(or method) is a function which is defined within function fanIn
	// this way is used when we want to define function within a function
	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}

// here we are transferring data from a function(generator function) to a stream
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	//func repeatFunc[T any](done <-chan int, fn func() T) <-chan T {

	stream := make(chan T)

	// forking a routine
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	// we are just returning the stream
	return stream
}

// here we are transferring data from one stream(stream) to another stream(taken)
func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()

	return taken
}

func randNumFetcher() int {
	return rand.Intn(50 * 1000000)
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randInt := <-randIntStream:
				if isPrime(randInt) {
					primes <- randInt
				}
			}
		}
	}()

	return primes
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	generator_stream := repeatFunc(done, randNumFetcher)
	// uncomment to test generator_stream
	// for rando := range generator_stream {
	// 	fmt.Println(rando)
	// }

	// 1) naive way
	// prime_stream := primeFinder(done, generator_stream)
	// take_stream := take(done, prime_stream, 10)
	// for rando := range take_stream {
	//   fmt.Println(rando)
	// }

	// 2) fanout - fanin way
	// fanout
	CPUCount := runtime.NumCPU()
	fmt.Println("available cpu", CPUCount)
	// declaring array of int channels
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, generator_stream)
	}

	// fanin
	fanned_in_stream := fanIn(done, primeFinderChannels...)
	take_stream := take(done, fanned_in_stream, 10)
	for rando := range take_stream {
		fmt.Println(rando)
	}

	fmt.Println(time.Since(start))
}
