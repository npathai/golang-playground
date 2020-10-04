package main

import (
	"../concurrency/book"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(6) + 1
		// Letting the Wait group know that we are creating 2 go routines
		wg.Add(2)
		// adding go keyword in front of the function makes it go routine, or like a Runnable.
		// Anonymous function and immediate call
		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
			// Task letting the wait group know that it's done
			wg.Done()
		}(id, wg)

		go func(id int, wg *sync.WaitGroup) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg)
	}

	// Waiting for all the routines to finish before exiting the main routine
	wg.Wait()
}

func queryCache(id int) (book.Book, bool) {
	book, ok := cache[id]
	return book, ok
}

func queryDatabase(id int) (book.Book, bool) {
	// Simulate lag
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return book.Book{}, false
}
