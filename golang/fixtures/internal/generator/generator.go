package generator

import (
	databaseAuthor "fixtures/internal/database/author"
	databaseBook "fixtures/internal/database/book"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func New() *Tool {
	return &Tool{}
}

type Tool struct {
}

func (t *Tool) CreateUUID() uuid.UUID {
	return uuid.New()
}

func (t *Tool) CreateBook() *databaseBook.Book {
	return &databaseBook.Book{
		ID:        t.CreateUUID(),
		Title:     gofakeit.BookTitle(),
		AuthorID:  t.CreateUUID(),
		IssueYear: int32(gofakeit.Year()),
	}
}

func (t *Tool) CreateAuthor() *databaseAuthor.Author {
	return &databaseAuthor.Author{
		ID:       t.CreateUUID(),
		Name:     gofakeit.Name(),
		HasBooks: false,
	}
}
