package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// Creates unbuffered channel
	//ch := make(chan int)

	// Creating a unbuffered channel
	ch := make(chan int, 2)

	wg.Add(3)

	// Receiver routine
	go func(ch chan int, wg *sync.WaitGroup) {
		// <-ch is for receiving value from the channel
		// This is blocking call. The routine is blocked on channel to receive a message
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// Sender routine
	go func(ch chan int, wg *sync.WaitGroup) {
		// ch <- int is for writing to the channel
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch, wg)

	// Sender routine which receives send only channel. Notice ch <- syntax
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 100
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
