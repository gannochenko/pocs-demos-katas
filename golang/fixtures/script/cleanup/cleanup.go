package main

import (
	"fixtures/internal/util/db"
)

func main() {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	connection.Exec("DELETE FROM authors")
	connection.Exec("DELETE FROM books")
}
