package book

import (
	"context"

	"gorm.io/gorm"

	"loggingerrorhandling/internal/domain/database/book"
	"loggingerrorhandling/internal/syserr"
)

const (
	TableName = "books"
	PageSize  = 5
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) GetBooks(ctx context.Context, filter string, page int32) (books []*book.Book, err error) {
	runner := r.Session.Table(TableName)
	if filter != "" {
		runner = runner.Where("title like ?", filter)
	}
	runner.Offset(int(page * PageSize)).Limit(PageSize).Find(&books)
	return books, nil
}

func (r *Repository) GetBookCount(ctx context.Context, filter string) (count int64, err error) {
	return 0, syserr.NewBadInput(
		"could not retrieve book count",
		syserr.F("term", filter),
		syserr.F("foo", "bar"),
	)

	runner := r.Session.Table(TableName)
	if filter != "" {
		runner = runner.Where("title like ?", filter)
	}
	runner.Select("id").Count(&count)
	return count, nil
}
