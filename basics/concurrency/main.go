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
	dbChan := make(chan book.Book)
	cacheChan := make(chan book.Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(6) + 1
		fmt.Printf("Id: %v\n", id)
		// Letting the Wait group know that we are creating 2 go routines
		wg.Add(2)
		// adding go keyword in front of the function makes it go routine, or like a Runnable.
		// Anonymous function and immediate call
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan <- book.Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			// Task letting the wait group know that it's done
			wg.Done()
		}(id, wg, m, cacheChan)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan <- book.Book) {
			if b, ok := queryDatabase(id, m); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, m, dbChan)

		// Create one go routine per query to receive the results
		go func(cacheChan, dbChan <- chan book.Book) {
			select {
			case b:= <- cacheChan:
				fmt.Println("From cache")
				fmt.Println(b)
				// When get the cache hit, we need to drain the db channel
				<- dbChan
			case b:= <- dbChan:
				fmt.Println("From database")
				fmt.Println(b)
			}
		}(cacheChan, dbChan)

		time.Sleep(150 * time.Millisecond)
	}

	// Waiting for all the routines to finish before exiting the main routine
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (book.Book, bool) {
	// Getting read lock to cache (shared memory)
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (book.Book, bool) {
	// Simulate lag
	time.Sleep(100 * time.Millisecond)
	for _, b := range book.Books {
		if b.ID == id {
			return b, true
		}
	}
	return book.Book{}, false
}
