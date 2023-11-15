package book

import (
	"fixtures/internal/domain/book"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title     string    `db:"title"`
	AuthorID  uuid.UUID `db:"author_id" gorm:"type:uuid"`
	IssueYear int32     `db:"issue_year"`
}

func (b *Book) ToDomain() *book.Book {
	return &book.Book{
		ID:        b.ID.String(),
		Title:     b.Title,
		AuthorID:  b.AuthorID.String(),
		IssueYear: b.IssueYear,
	}
}

type ListParametersFilter struct {
	Title    *string
	AuthorID *string
}

type ListParameters struct {
	Filter *ListParametersFilter
	Page   int32
}
