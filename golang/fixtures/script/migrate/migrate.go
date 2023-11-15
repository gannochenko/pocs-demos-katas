package main

import (
	"fixtures/internal/database/author"
	"fixtures/internal/database/book"
	"fixtures/internal/util/db"
)

func main() {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	connection.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Migrate the schema
	err = connection.AutoMigrate(&book.Book{})
	if err != nil {
		panic(err)
	}

	err = connection.AutoMigrate(&author.Author{})
	if err != nil {
		panic(err)
	}
}
