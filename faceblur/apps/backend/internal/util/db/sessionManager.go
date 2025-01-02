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

func (m *SessionManager) Begin(handle interfaces.SessionHandle) (interfaces.SessionHandle, error) {
	if handle != nil {
		return handle, nil
	}

	tx := m.session.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &SessionHandle{
		tx: tx,
	}, nil
}

func (m *SessionManager) Commit(handle interfaces.SessionHandle) error {
	if handle != nil {
		tx := handle.GetTx()
		tx.Commit()
		return tx.Error
	}

	return nil
}

func (m *SessionManager) RollbackUnlessCommitted(handle interfaces.SessionHandle) error {
	if r := recover(); r != nil { // ?? will this even work?
		if handle != nil {
			tx := handle.GetTx()
			tx.Rollback()
			return tx.Error
		}
	}

	return nil
}

type SessionHandle struct {
	tx *gorm.DB
}

func (h *SessionHandle) GetTx() *gorm.DB {
	return h.tx
}
