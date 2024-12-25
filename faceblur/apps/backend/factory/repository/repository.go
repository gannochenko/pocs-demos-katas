package repository

import (
	"gorm.io/gorm"
)

type Factory struct {
	session *gorm.DB

	//bookRepository *book.Repository
}

func NewRepositoryFactory(session *gorm.DB) *Factory {
	return &Factory{
		session: session,
	}
}

//
//func (m *Manager) GetBookRepository() interfaces.BookRepository {
//	if m.bookRepository == nil {
//		m.bookRepository = book.New(m.session)
//	}
//
//	return m.bookRepository
//}
