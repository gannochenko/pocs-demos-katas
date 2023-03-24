package book

import (
	"gorm.io/gorm"
	"levelsgorm/internal/domain/business/book"
)

type Book struct {
	gorm.Model
	ID        string `db:"id"`
	Title     string `db:"title"`
	Author    string `db:"author"`
	IssueYear int32  `db:"issue_year"`
}

func (b *Book) ToBusiness() *book.Book {
	return &book.Book{
		ID:        b.ID,
		Title:     b.Title,
		Author:    b.Author,
		IssueYear: b.IssueYear,
	}
}
