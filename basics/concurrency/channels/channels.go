package main

import (
	"fmt"
	"sync"
	"time"
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
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		// After receiving 3 messages that the senders send. Post the channel is closed the the reader will start receiving '0' value
		// We need some mechanism on reader's side to know when a channel is closed
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
		// Allow other writer to write, so that we can send and close the channel
		time.Sleep(1 * time.Second)

		ch <- 100
		// Closing the channel on the sender side. No new messages can be published on channel now
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
