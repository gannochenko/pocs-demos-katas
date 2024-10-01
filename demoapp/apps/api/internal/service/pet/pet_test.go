package pet_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"api/internal/domain"
	"api/internal/factory"
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
	serviceFactory := factory.MakeServiceFactory(session)
	dataGenerator := test.NewGenerator()
	dataBuilder := test.NewBuilder(session)
	ctx := context.Background()

	type setup struct {
		request *domain.ListPetsRequest
		ids     []string
	}
	type verify struct {
		err      error
		response *domain.ListPetsResponse
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should return list of pets filtered by ID": {
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
					request: &domain.ListPetsRequest{
						IDs: []string{pet1.ID.String(), pet2.ID.String()},
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				var pet1 *domain.Pet
				var pet2 *domain.Pet
				for _, pet := range verify.response.Pets {
					if pet.ID == setup.ids[0] {
						pet1 = pet
					}
					if pet.ID == setup.ids[1] {
						pet2 = pet
					}
				}

				assert.NotNil(t, pet1)
				assert.Equal(t, domain.PetStatusAvailable, pet1.Status)
				assert.True(t, len(pet1.Name) > 0)

				assert.NotNil(t, pet2)
				assert.Equal(t, domain.PetStatusAvailable, pet2.Status)
				assert.True(t, len(pet2.Name) > 0)

				assert.Equal(t, int32(50), verify.response.Pagination.PageSize)
				assert.Equal(t, int64(1), verify.response.Pagination.PageCount)
				assert.Equal(t, int64(2), verify.response.Pagination.Total)
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			response, err := serviceFactory.GetPetService().ListPets(ctx, setup.request)
			testCase.verifyFunc(t, setup, &verify{
				err:      err,
				response: response,
			})
		})
	}
}
