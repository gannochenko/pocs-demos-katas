package author

import (
	"context"

	"gorm.io/gorm"
)

const (
	TableName = "authors"
)

type Repository struct {
	Session *gorm.DB
}

func (r *Repository) RefreshHasBooksFlag(ctx context.Context, condition interface{}) (err error) {
	//runner := r.Session.Table(TableName)

	return nil
}
