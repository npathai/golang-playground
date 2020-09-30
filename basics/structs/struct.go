package main

import "fmt"

func main() {
	type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u user
	// All the values inside the struct are initialized to their default values
	// So ID -> 0, firstName and lastName will be initialized to blank values
	fmt.Println(u)

	var u1 user
	u1.ID = 123
	u1.FirstName = "ABC"
	u1.LastName = "XYZ"
	fmt.Println(u1)

	u2 := user {ID: 321,
		FirstName: "CDE",
		LastName: "EFG",
	}
	fmt.Println(u2)


}
