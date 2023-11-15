package database_writer

import (
	databaseAuthor "fixtures/internal/database/author"
	databaseBook "fixtures/internal/database/book"
	"gorm.io/gorm"
)

type Writer struct {
	session *gorm.DB

	books   []*databaseBook.Book
	authors []*databaseAuthor.Author
}

func New(session *gorm.DB) *Writer {
	writer := &Writer{
		session: session,
	}
	writer.resetArrays()

	return writer
}

func (w *Writer) AddBook(book *databaseBook.Book) {
	w.books = append(w.books, book)
}

func (w *Writer) AddAuthor(author *databaseAuthor.Author) {
	w.authors = append(w.authors, author)
}

func (w *Writer) Dump() error {
	for _, author := range w.authors {
		w.session.Create(author)
	}

	for _, book := range w.books {
		w.session.Create(book)
	}

	return nil
}

func (w *Writer) Reset() {
	w.resetArrays()
}

func (w *Writer) resetArrays() {
	w.books = make([]*databaseBook.Book, 0)
	w.authors = make([]*databaseAuthor.Author, 0)
}
