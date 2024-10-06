package category_test

import (
	"context"
	"os"
	"testing"

	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"

	"api/internal/dto"
	"api/internal/factory"
	"api/internal/service/config"
	"api/internal/util/db"
	"api/test"
)

var (
	session *gorm.DB
)

func TestMain(m *testing.M) {
	var err error
	session, err = db.Connect(os.Getenv("POSTGRES_DB_DSN"))
	if err != nil {
		panic(err)
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestListCategories(t *testing.T) {
	configService := config.NewConfigService()

	serviceFactory := factory.MakeServiceFactory(session, configService)
	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)
	ctx := context.Background()

	type setup struct {
		parameters *dto.ListCategoriesParameters
		ids        []string
	}
	type verify struct {
		err    error
		result []*dto.Category
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return list of pets": {
			setupFunc: func(t *testing.T) *setup {
				category1 := dataGenerator.CreateCategory()
				category2 := dataGenerator.CreateCategory()

				assert.NoError(t, dataBuilder.
					Reset().
					AddCategories(category1, category2).
					Submit(),
				)

				return &setup{
					ids: []string{category1.ID.String(), category2.ID.String()},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				var category1 *dto.Category
				var category2 *dto.Category
				for _, category := range verify.result {
					if category.ID.String() == setup.ids[0] {
						category1 = category
					}
					if category.ID.String() == setup.ids[1] {
						category2 = category
					}
				}

				assert.NotNil(t, category1)
				assert.NotNil(t, category2)
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			result, err := serviceFactory.GetRepositoryFactory().GetCategoryRepository().ListCategories(ctx, nil, setup.parameters)
			testCase.verifyFunc(t, setup, &verify{
				err:    err,
				result: result,
			})
		})
	}
}
