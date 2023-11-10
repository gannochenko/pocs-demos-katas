package author

import (
	"gorm.io/gorm"
)

const (
	TableName = "authors"
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) RefreshHasBooksFlag() (err error) {
	//runner := r.Session.Table(TableName)

	return nil
}
