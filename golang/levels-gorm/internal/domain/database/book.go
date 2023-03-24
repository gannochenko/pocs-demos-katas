package database

import "levelsgorm/internal/domain/business"

type Book struct {
	ID        string `db:"id"`
	Title     string `db:"title"`
	Author    string `db:"author"`
	IssueYear int32  `db:"issue_year"`
}

func (b *Book) ToBusiness() *business.Book {
	return &business.Book{
		ID:        b.ID,
		Title:     b.Title,
		Author:    b.Author,
		IssueYear: b.IssueYear,
	}
}
