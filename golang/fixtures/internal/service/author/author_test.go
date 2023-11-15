package author_test

import (
	"context"
	"os"
	"testing"

	authorDomain "fixtures/internal/domain/author"
	authorRepository "fixtures/internal/repository/author"
	bookRepository "fixtures/internal/repository/book"
	authorService "fixtures/internal/service/author"
	"fixtures/internal/util/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	ctx     context.Context
	session *gorm.DB
)

func TestMain(m *testing.M) {
	var err error
	session, err = db.Connect()
	if err != nil {
		panic(err)
	}

	ctx = context.Background()

	execCode := m.Run()
	os.Exit(execCode)
}

func TestGetBooksAndAuthor(t *testing.T) {
	type setup struct {
		authorID string
	}
	type verify struct {
		err    error
		result *authorDomain.GetBooksAndAuthorResult
	}

	booksRepo := &bookRepository.Repository{
		Session: session,
	}
	authorsRepo := &authorRepository.Repository{
		Session: session,
	}
	authorSvc := authorService.New(booksRepo, authorsRepo)

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return a result": {
			setupFunc: func(t *testing.T) *setup {
				return &setup{}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := tc.setupFunc(t)

			result, err := authorSvc.GetBooksAndAuthor(ctx, setup.authorID)

			tc.verifyFunc(t, setup, &verify{
				err:    err,
				result: result,
			})
		})
	}
}
