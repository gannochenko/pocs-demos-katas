package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"api/internal/api"
	"api/internal/domain"
	"api/test"
)

func TestToDomain(t *testing.T) {
	type setup struct {
		item *api.Pet
	}
	type verify struct {
		item *domain.Pet
		err  error
	}

	type setupFunc func(t *testing.T) *setup
	type verifyFunc func(t *testing.T, setup *setup, verify *verify)

	dataGenerator := test.NewGenerator()

	testCases := map[string]struct {
		setupFunc  setupFunc
		verifyFunc verifyFunc
	}{
		"Should convert api to domain": {
			setupFunc: func(t *testing.T) *setup {
				pet := dataGenerator.CreatePet()

				return &setup{
					item: &api.Pet{
						ID:        pet.ID.String(),
						Name:      pet.Name,
						Status:    string(pet.Status),
						PhotoUrls: pet.PhotoUrls,
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			setup := testCase.setupFunc(t)

			item, err := setup.item.ToDomain()

			testCase.verifyFunc(t, setup, &verify{
				item: item,
				err:  err,
			})
		})
	}
}
