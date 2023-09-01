package repository

import (
	"dependencymanager/internal/interfaces"
	"dependencymanager/internal/repository/book"
	"gorm.io/gorm"
)

type Manager struct {
	session *gorm.DB

	bookRepository *book.Repository
}

func New(session *gorm.DB) *Manager {
	return &Manager{
		session: session,
	}
}

func (m *Manager) GetBookRepository() interfaces.BookRepository {
	if m.bookRepository == nil {
		m.bookRepository = book.New(m.session)
	}

	return m.bookRepository
}
