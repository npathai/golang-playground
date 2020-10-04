package main

import "fmt"

func main() {

	// Comparison operator
	b1 := 1 == 1
	b2 := 1 == 2
	fmt.Printf("%v %T \n", b1, b1)
	fmt.Printf("%v %T \n", b2, b2)

	// Bitwise operators
	a := 10
	b := 3

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	a1 := 8
	fmt.Println(a1 << 3)
	fmt.Println(a1 >> 3)
}
