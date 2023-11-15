package author

import (
	authorDomain "fixtures/internal/domain/author"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string    `db:"name"`
	HasBooks bool      `db:"has_books"`
}

func (b *Author) ToDomain() *authorDomain.Author {
	return &authorDomain.Author{
		ID:   b.ID.String(),
		Name: b.Name,
	}
}

type ListParametersFilter struct {
	ID *string
}

type ListParameters struct {
	Filter *ListParametersFilter
	Page   int32
}
