package main

import "fmt"

func main() {
	// This is useful if we want to initialize the variable from different place
	var i int
	i = 123
	fmt.Println(i)

	// This is useful shorthand if we want to initialize the variable with declaration
	var f float32 = 123.12
	fmt.Println(f)

	// implicit initialization syntax
	firstName := "Narendra"
	fmt.Println(firstName)

	// boolean data type
	isAdult := true
	fmt.Println(isAdult)

	// Complex data type
	c := complex(3, 4)
	fmt.Println(c)

	// Pull out the real and imaginary component of complex data type
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// Printing the type of the variable
	truthy := false
	fmt.Printf("%v, %T\n\n", truthy, truthy)
}
