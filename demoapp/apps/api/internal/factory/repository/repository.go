package repository

import (
	"gorm.io/gorm"
)

type Factory struct {
	session *gorm.DB

	//bookRepository *book.Repository
}

func New(session *gorm.DB) *Factory {
	return &Factory{
		session: session,
	}
}

//func (m *Factory) GetBookRepository() interfaces.BookRepository {
//	if m.bookRepository == nil {
//		m.bookRepository = book.New(m.session)
//	}
//
//	return m.bookRepository
//}
