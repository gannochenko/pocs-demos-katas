package main

import (
	"levelsgorm/internal/domain/database/book"
	"levelsgorm/internal/util/db"
)

func main() {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	connection.Create(&book.Book{
		Title:     "1984",
		Author:    "Orwell",
		IssueYear: 1949,
	})
	connection.Create(&book.Book{
		Title:     "Animal Farm",
		Author:    "Orwell",
		IssueYear: 1945,
	})
	connection.Create(&book.Book{
		Title:     "Fahrenheit 451",
		Author:    "Bradbury",
		IssueYear: 1953,
	})
	connection.Create(&book.Book{
		Title:     "Do Androids Dream of Electric Sheep?",
		Author:    "Dick",
		IssueYear: 1968,
	})
	connection.Create(&book.Book{
		Title:     "The Goblin Reservation",
		Author:    "Simak",
		IssueYear: 1968,
	})
	connection.Create(&book.Book{
		Title:     "Childhood's End",
		Author:    "Clarke",
		IssueYear: 1968,
	})
}
