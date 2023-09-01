package service

import (
	"dependencymanager/internal/interfaces"
	"dependencymanager/internal/manager/repository"
	bookService "dependencymanager/internal/service/book"
	"gorm.io/gorm"
)

type Manager struct {
	session           *gorm.DB
	repositoryManager *repository.Manager

	bookService *bookService.Service
}

func New(session *gorm.DB, repositoryManager *repository.Manager) *Manager {
	return &Manager{
		session:           session,
		repositoryManager: repositoryManager,
	}
}

func (m *Manager) GetBookService() interfaces.BookService {
	if m.bookService == nil {
		m.bookService = bookService.New(m.session, m.repositoryManager.GetBookRepository())
	}

	return m.bookService
}
