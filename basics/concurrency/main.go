package main

import (
	"../concurrency/book"
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]book.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(6) + 1
		// adding go keyword in front of the function makes it go routine, or like a Runnable.
		// Anonymous function and immediate call
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println("From cache")
				fmt.Println(b)
			}
		}(id)

		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				fmt.Println("From database")
				fmt.Println(b)
			}
		}(id)

		time.Sleep(150 * time.Millisecond)
	}

	// Wait for the Go routines to finish
	time.Sleep(2 * time.Second)
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
