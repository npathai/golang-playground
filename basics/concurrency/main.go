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
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(6) + 1
		// Letting the Wait group know that we are creating 2 go routines
		wg.Add(2)
		// adding go keyword in front of the function makes it go routine, or like a Runnable.
		// Anonymous function and immediate call
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
			// Task letting the wait group know that it's done
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("From database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)
	}

	// Waiting for all the routines to finish before exiting the main routine
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	// Getting read lock to cache (shared memory)
	m.RLock()
	book, ok := cache[id]
	m.RUnlock()
	return book, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	// Simulate lag
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return book.Book{}, false
}
