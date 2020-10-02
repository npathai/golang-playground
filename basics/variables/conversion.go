package main

import (
	"strconv"
)

func main() {
	i := 1

	// Converting i to a float 32
	f := float32(i)
	println(f)

	// Explicit conversion - Can't convert from floats to ints due to loss of data. Need explicit conversion
	f1 := 42.5
	var i2 int = int(f1)
	println(i2)

	// Converting from integer to string

	// Doesn't work this way - Will print the unicode value represented by integer
	i3 := 42
	println(string(i3))

	// To convert integer to string, we need help of strconv package
	str := strconv.Itoa(i3)
	println(str)

	// Converting from string to int
	parsedInt, _ := strconv.Atoi(str)
	println(parsedInt)
}
