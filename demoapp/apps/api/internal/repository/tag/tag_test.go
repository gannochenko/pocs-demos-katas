package tag_test

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

func TestListTags(t *testing.T) {
	configService := config.NewConfigService()

	serviceFactory := factory.MakeServiceFactory(session, configService)
	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)
	ctx := context.Background()

	type setup struct {
		parameters *dto.ListTagsParameters
		ids        []string
	}
	type verify struct {
		err    error
		result []*dto.Tag
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return list of pets": {
			setupFunc: func(t *testing.T) *setup {
				tag1 := dataGenerator.CreateTag()
				tag2 := dataGenerator.CreateTag()

				assert.NoError(t, dataBuilder.
					Reset().
					AddTags(tag1, tag2).
					Submit(),
				)

				return &setup{
					ids: []string{tag1.ID.String(), tag2.ID.String()},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				var tag1 *dto.Tag
				var tag2 *dto.Tag
				for _, tag := range verify.result {
					if tag.ID.String() == setup.ids[0] {
						tag1 = tag
					}
					if tag.ID.String() == setup.ids[1] {
						tag2 = tag
					}
				}

				assert.NotNil(t, tag1)
				assert.NotNil(t, tag2)
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			result, err := serviceFactory.GetRepositoryFactory().GetTagRepository().ListTags(ctx, nil, setup.parameters)
			testCase.verifyFunc(t, setup, &verify{
				err:    err,
				result: result,
			})
		})
	}
}
