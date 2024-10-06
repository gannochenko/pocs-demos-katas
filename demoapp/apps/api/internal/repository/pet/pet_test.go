package pet_test

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

func TestListPets(t *testing.T) {
	configService := config.NewConfigService()

	serviceFactory := factory.MakeServiceFactory(session, configService)
	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)
	ctx := context.Background()

	type setup struct {
		parameters *dto.ListPetParameters
		ids        []string
		category   *dto.Category
	}
	type verify struct {
		err    error
		result []*dto.Pet
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return list of pets": {
			setupFunc: func(t *testing.T) *setup {
				pet1 := dataGenerator.CreatePet()
				pet2 := dataGenerator.CreatePet()

				assert.NoError(t, dataBuilder.
					Reset().
					AddPets(pet1, pet2).
					Submit(),
				)

				return &setup{
					ids: []string{pet1.ID.String(), pet2.ID.String()},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				var pet1 *dto.Pet
				var pet2 *dto.Pet
				for _, pet := range verify.result {
					if pet.ID.String() == setup.ids[0] {
						pet1 = pet
					}
					if pet.ID.String() == setup.ids[1] {
						pet2 = pet
					}
				}

				assert.NotNil(t, pet1)
				assert.NotNil(t, pet2)
			},
		},
		"Should select categories": {
			setupFunc: func(t *testing.T) *setup {
				pet1 := dataGenerator.CreatePet()
				category1 := dataGenerator.CreateCategory()
				pet1.CategoryID = &category1.ID

				assert.NoError(t, dataBuilder.
					Reset().
					AddPets(pet1).
					AddCategories(category1).
					Submit(),
				)

				return &setup{
					ids:      []string{pet1.ID.String()},
					category: category1,
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				var pet1 *dto.Pet
				for _, pet := range verify.result {
					if pet.ID.String() == setup.ids[0] {
						pet1 = pet
					}
				}

				assert.NotNil(t, pet1)
				assert.Equal(t, setup.category.ID.String(), pet1.Category.ID)
				assert.Equal(t, setup.category.Name, pet1.Category.Name)
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			result, err := serviceFactory.GetRepositoryFactory().GetPetRepository().ListPets(ctx, nil, setup.parameters)
			testCase.verifyFunc(t, setup, &verify{
				err:    err,
				result: result,
			})
		})
	}
}
