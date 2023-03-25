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
}
