package main

import (
	"errors"
	"fmt"
)

func main() {
	func1()

	arg := 123
	arg2 := "Hello world!"
	func2(arg, arg2)

	fmt.Println(add(1, 2))
	sum := add(3, 4)
	fmt.Println(sum)

	err := start(3000)
	fmt.Println(err)

	output, err := stop()
	fmt.Println(output, err)

	// We can use _ to create write only variable, which we don't want to use and ignore
	_, err1 := stop()
	fmt.Println(err1)
}

func func1() {
	fmt.Println("func1 called")
}

func func2(arg int, message string) {
	fmt.Println(arg, message)
}

func add(arg1, arg2 int) int {
	return arg1 + arg2
}

func start(port int) error {
	fmt.Println("Started on port: ", port)
	return errors.New("something went wrong")
}

func stop() (int, error) {
	return 12, nil
}