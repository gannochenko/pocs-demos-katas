package test

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"

	"backend/internal/database"
)

func NewGenerator() *Generator {
	return &Generator{}
}

type Generator struct {
}

func (g *Generator) CreateUUID() uuid.UUID {
	return uuid.New()
}

func (g *Generator) CreateImage() *database.Image {
	return &database.Image{
		ID: g.CreateUUID(),
		//URL:         lo.ToPtr(gofakeit.URL()),
		OriginalURL: gofakeit.URL(),
		CreatedBy:   g.CreateUUID(),
		CreatedAt:   g.CreateUTCTime(),
		UpdatedAt:   g.CreateUTCTime(),
		IsProcessed: false,
	}
}

func (g *Generator) CreateUser() *database.User {
	return &database.User{
		ID:        g.CreateUUID(),
		Sup:       fmt.Sprintf("auth0:%d", g.CreatePositiveInt32()),
		Email:     gofakeit.Email(),
		CreatedAt: g.CreateUTCTime(),
		UpdatedAt: g.CreateUTCTime(),
	}
}

func (g *Generator) CreatePositiveInt32() int32 {
	return int32(gofakeit.Uint16() + 1)
}

func (g *Generator) CreateUTCTime() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}
