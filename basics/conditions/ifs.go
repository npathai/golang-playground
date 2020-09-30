package main

type User struct {
	ID int
	FirstName string
	LastName string
}

func main() {
	u1 := User{
		ID: 1,
		FirstName: "Ada",
		LastName: "Lovelace",
	}
	u2 := User {
		ID: 2,
		FirstName: "Donald",
		LastName: "Knuth",
	}

	if u1 == u2 {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		println("Similar users")
	} else {
		println("Different users")
	}
}
