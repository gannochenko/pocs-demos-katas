package interfaces

import "gorm.io/gorm"

type SessionHandle interface {
	GetTx() *gorm.DB
}

type SessionManager interface {
	Begin(handle SessionHandle) (SessionHandle, error)
	RollbackUnlessCommitted(handle SessionHandle) error
	Commit(handle SessionHandle) error
}
