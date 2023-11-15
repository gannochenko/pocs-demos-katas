package author

import (
	bookDomain "fixtures/internal/domain/book"
)

type Author struct {
	ID   string
	Name string
}

type GetBooksAndAuthorResult struct {
	Books  []*bookDomain.Book
	Author *Author
}
