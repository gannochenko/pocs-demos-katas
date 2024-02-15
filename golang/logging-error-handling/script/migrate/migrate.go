package main

import (
	"loggingerrorhandling/internal/domain/database/book"
	"loggingerrorhandling/internal/util/db"
)

func main() {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	connection.AutoMigrate(&book.Book{})
}
