package test

import (
	"gorm.io/gorm"

	"backend/internal/database"
)

type Builder struct {
	session *gorm.DB

	images []*database.Image
	users  []*database.User
}

func NewBuilder(session *gorm.DB) *Builder {
	writer := &Builder{
		session: session,
	}
	writer.resetArrays()

	return writer
}

func (b *Builder) AddImage(image *database.Image) {
	b.images = append(b.images, image)
}

func (b *Builder) AddImages(images ...*database.Image) *Builder {
	for _, image := range images {
		b.AddImage(image)
	}

	return b
}

func (b *Builder) AddUser(user *database.User) {
	b.users = append(b.users, user)
}

func (b *Builder) AddUsers(users ...*database.User) *Builder {
	for _, user := range users {
		b.AddUser(user)
	}

	return b
}

func (b *Builder) Submit() error {
	for _, user := range b.users {
		res := b.session.Create(user)
		if res.Error != nil {
			return res.Error
		}
	}

	for _, image := range b.images {
		res := b.session.Create(image)
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func (b *Builder) Reset() *Builder {
	b.resetArrays()

	return b
}

func (b *Builder) TruncateAll() *Builder {
	b.session.
		Exec("TRUNCATE TABLE images").
		Exec("TRUNCATE TABLE users")

	return b
}

func (b *Builder) SelectImages(filter string) (result []*database.Image, err error) {
	queryResult := b.session.Table("images").Where(filter).Find(&result)
	if queryResult.Error != nil {
		return nil, queryResult.Error
	}
	return result, nil
}

func (b *Builder) resetArrays() {
	b.images = make([]*database.Image, 0)
	b.users = make([]*database.User, 0)
}
