package main

import "fmt"

// Arrays are fixed in size and cannot be resized. It is compile time error to try to access element past the size
func main() {
	// Long form of creating the array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Implicit Initialization
	arr2 := [3]int {1, 2, 3}
	fmt.Println(arr2)
}
