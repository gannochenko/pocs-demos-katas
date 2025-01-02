package db

import "gorm.io/gorm"

type SessionManager struct {
	session *gorm.DB
}
