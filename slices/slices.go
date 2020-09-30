package main

import "fmt"

// They are based on top of Array. We can resize a slice, so they are dynamically sized
func main() {
	arr := [3]int{1, 2, 3}

	// Creating a slice of whole array
	slice := arr[:]

	fmt.Println(arr, slice)

	arr[0] = 12
	slice[1] = 23
	// Changing either the arr or slice will reflect in both, because slice is pointing to the data that array is keeping
	fmt.Println(arr, slice)

	// Creating a slice without an array, with dynamic sizing
	// Underlying array is going to be managed by Go runtime
	dynamicSlice := []int{1, 2, 3}
	fmt.Println(dynamicSlice)

	// If the size exhausts while appending the append function will automatically increment the size and return new array instance
	dynamicSlice = append(dynamicSlice, 4, 5, 6)
	fmt.Println(dynamicSlice)

	// Creating a subset slice from an existing slice
	// From inclusive, to exclusive like most of such operations in other languages
	slice2 := slice[2:]
	slice3 := slice[:2]
	slice4 := slice[1:2]
	fmt.Println(slice2, slice3, slice4)
}
