package main

import "fmt"

// This is an constant at the package level. It is accessible throughout the whole package
const pi = 3.1415

// This is constant block just like import block
const (
	// Everytime iota is used it increments its value by 1
	first = iota
	second = iota
)

const (
	// In every constant block iota is reset to 0
	// That's why value of first and another both will be 0
	another = iota
)

const (
	firstOne = iota + 5
	// If we don't explicitly initialize the value, then it will assume we need to increment the value of previous
	// assignment and secondOne will get the value of 6
	secondOne
)

func main() {
	// Due to iota, first will be equal to 0 and second will be equal to 1
	fmt.Println(first, second)
	fmt.Println(another)
	fmt.Println(firstOne, secondOne)
}
