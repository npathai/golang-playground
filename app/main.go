package main

import (
	"./models"
	"fmt"
)
func main() {
	user := models.User {
		ID : 1,
		FirstName : "ABC",
		LastName : "CDE",
	}
	fmt.Println(user)
}
