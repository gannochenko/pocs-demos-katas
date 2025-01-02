package db

import (
	"gorm.io/gorm"

	"backend/interfaces"
)

type SessionManager struct {
	session *gorm.DB
}

func NewSessionManager(session *gorm.DB) *SessionManager {
	return &SessionManager{
		session: session,
	}
}

func (m *SessionManager) Begin() (interfaces.SessionHandle, error) {
	return &SessionHandle{
		tx: nil,
	}, nil
}

type SessionHandle struct {
	tx *gorm.DB
}

func (h *SessionHandle) Commit() {

}

func (h *SessionHandle) Rollback() {

}
