package main

import "fmt"

func main() {
	// Can't change the value of the constant once it's initialized
	// We are not allowed to initialize the constant at some other place, must be initialized during declaration

	// Implicit declaration syntax
	const pi = 3.1415
	fmt.Println(pi)

	// Constants must be COMPILE time constants, this line will not work
	//const dynamic = value()

	const i = 3
	// Here the compiler will consider const i to be an integer
	fmt.Println(i + 2)

	// Here the compiler will consider const i to be a floating point number
	fmt.Println(i + 1.23)


	// Explicit declaration syntax
	const restrictedType int = 1
	// Will not allow to add floating point to an integer, we need explicit conversion operation
	//fmt.Println(restrictedType + 1.3)

	// Converted to explicit float32 using conversion operation
	fmt.Println(float32(restrictedType) + 1.3)
}