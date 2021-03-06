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

		// Using the comma ok syntax to check if legitimate 0 was written by sender or the channel is closed
		if msg, ok := <- ch; ok {
			fmt.Println(msg)
		}
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


	wg1 := &sync.WaitGroup{}
	ch1 := make(chan int)

	wg1.Add(2)

	// Writer routine
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch1, wg1)

	// Reader routine
	go func(ch <-chan int, wg *sync.WaitGroup) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch1, wg1)

	wg1.Wait()
}