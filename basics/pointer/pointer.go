package main

import "fmt"

func main() {
	// Example of declaring a pointer which is uninitialized. It will have default value of nil
	var firstName *string
	fmt.Println(firstName)

	// De-referencing the pointer and assigning it value
	// This won't work because pointer is nil, go doesn't allow assigning value to nil pointer
	//*firstName = "ABC"
	fmt.Println(firstName)

	// So to make the above code work we will have to set aside memory for that pointer of data type it will hold
	var lastName *string = new(string)
	*lastName = "CDE"
	// This will work but it will print the address, as it is a pointer
	fmt.Println(lastName)

	// So to de-reference the pointer we need to use the same syntax *pointerName to get the value pointed by the pointer
	fmt.Println(*lastName)


	// Address of operator
	middleName := "XYZ"
	fmt.Println(middleName)

	ptr := &middleName
	fmt.Println(ptr, *ptr)

	// By changing the value of middle name the value of pointer does not change, the value does change.
	// Because middle name variable is still in the same location in the memory, we just changed the data that's stored there
	middleName = "PQR"
	fmt.Println(ptr, *ptr)
}
