package main

import (
	"hookspattern/internal/domain/database/author"
	"hookspattern/internal/domain/database/book"
	"hookspattern/internal/util/db"
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
