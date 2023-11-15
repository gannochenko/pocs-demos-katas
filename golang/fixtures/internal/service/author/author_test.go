package author_test

import (
	"context"
	"os"
	"testing"

	databaseWriter "fixtures/internal/database_writer"
	authorDomain "fixtures/internal/domain/author"
	"fixtures/internal/generator"
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

	gen := generator.New()
	writer := databaseWriter.New(session)

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return a result": {
			setupFunc: func(t *testing.T) *setup {
				writer.Reset()

				author1 := gen.CreateAuthor()

				book1 := gen.CreateBook()
				book1.AuthorID = author1.ID
				book2 := gen.CreateBook()
				book2.AuthorID = author1.ID
				book3 := gen.CreateBook()
				book3.AuthorID = author1.ID

				writer.AddAuthor(author1)
				writer.AddBook(book1)
				writer.AddBook(book2)
				writer.AddBook(book3)

				err := writer.Dump()
				assert.NoError(t, err)

				return &setup{
					authorID: author1.ID.String(),
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)
				assert.Equal(t, setup.authorID, verify.result.Author.ID)
				assert.Equal(t, 3, len(verify.result.Books))
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
