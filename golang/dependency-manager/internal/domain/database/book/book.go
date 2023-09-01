package book

import (
	"dependencymanager/internal/domain/business/book"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	IssueYear int32     `db:"issue_year"`
}

func (b *Book) ToBusiness() *book.Book {
	return &book.Book{
		ID:        b.ID.String(),
		Title:     b.Title,
		Author:    b.Author,
		IssueYear: b.IssueYear,
	}
}
