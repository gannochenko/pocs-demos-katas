package main

import (
	"github.com/google/uuid"
	"hookspattern/internal/database/author"
	"hookspattern/internal/database/book"
	"hookspattern/internal/util/db"
)

func main() {
	connection, err := db.Connect()
	if err != nil {
		panic(err)
	}

	orwell := &author.Author{
		ID:   uuid.New(),
		Name: "Orwell",
	}
	connection.Create(orwell)

	bradbury := &author.Author{
		ID:   uuid.New(),
		Name: "Bradbury",
	}
	connection.Create(bradbury)

	dick := &author.Author{
		ID:   uuid.New(),
		Name: "Dick",
	}
	connection.Create(dick)

	simak := &author.Author{
		ID:   uuid.New(),
		Name: "Simak",
	}
	connection.Create(simak)

	clarke := &author.Author{
		ID:   uuid.New(),
		Name: "Clarke",
	}
	connection.Create(clarke)

	marx := &author.Author{
		ID:   uuid.New(),
		Name: "Marx",
	}
	connection.Create(marx)

	//// ////////

	connection.Create(&book.Book{
		Title:     "1984",
		AuthorID:  orwell.ID,
		IssueYear: 1949,
	})
	connection.Create(&book.Book{
		Title:     "Animal Farm",
		AuthorID:  orwell.ID,
		IssueYear: 1945,
	})
	connection.Create(&book.Book{
		Title:     "Fahrenheit 451",
		AuthorID:  bradbury.ID,
		IssueYear: 1953,
	})
	connection.Create(&book.Book{
		Title:     "Do Androids Dream of Electric Sheep?",
		AuthorID:  dick.ID,
		IssueYear: 1968,
	})
	connection.Create(&book.Book{
		Title:     "The Goblin Reservation",
		AuthorID:  simak.ID,
		IssueYear: 1968,
	})
	connection.Create(&book.Book{
		Title:     "Childhood's End",
		AuthorID:  clarke.ID,
		IssueYear: 1968,
	})

	connection.Exec(`
		UPDATE
			authors
		SET has_books = bookinfo.book_count > 0
		FROM
			(
				SELECT
					authors.id AS author_id,
					count(books.*) AS book_count
				FROM
					authors
						LEFT JOIN books
								  ON books.author_id = authors.id
				GROUP BY
					authors.id
			) AS bookinfo
		WHERE
			authors.id = bookinfo.author_id
	`)
}
