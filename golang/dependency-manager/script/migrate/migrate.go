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

	// Migrate the schema
	connection.AutoMigrate(&book.Book{})
}
