package book

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t\t%v\n", b.Title, b.Author, b.YearPublished)
}

var Books = []Book{
	{
		1,
		"The Hitchhiker's Guide to the Galaxy",
		"Douglas Adams",
		1979,
	},

	{
		2,
		"The Hobbit",
		"J.R.R Tolkien",
		1937,
	},

	{
		3,
		"A Tale of Two Cities",
		"Charles Dickens",
		1859,
	},

	{
		4,
		"Les Miserables",
		"Victor Hugo",
		1862,
	},

	{
		5,
		"Harry Potter and the Philosopher's stone",
		"J.K. Rowling",
		1997,
	},

	{
		6,
		"I, Robot",
		"Issac Asimov",
		1950,
	},
}
