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
				tag1 := dataGenerator.CreateTag()
				tag2 := dataGenerator.CreateTag()
				category := dataGenerator.CreateCategory()

				return &setup{
					item: &api.Pet{
						ID:        pet.ID.String(),
						Name:      pet.Name,
						Status:    string(pet.Status),
						PhotoUrls: pet.PhotoUrls,
						Category: api.PetCategory{
							ID:   category.ID.String(),
							Name: category.Name,
						},
						Tags: []api.PetTag{
							{
								ID:   tag1.ID.String(),
								Name: tag1.Name,
							},
							{
								ID:   tag2.ID.String(),
								Name: tag2.Name,
							},
						},
					},
				}
			},
			verifyFunc: func(t *testing.T, setup *setup, verify *verify) {
				assert.NoError(t, verify.err)

				assert.Equal(t, setup.item.ID, verify.item.ID)
				assert.Equal(t, setup.item.Name, verify.item.Name)
				assert.Equal(t, setup.item.Status, string(verify.item.Status))
				assert.Equal(t, setup.item.PhotoUrls, verify.item.PhotoUrls)
				assert.Equal(t, setup.item.Category.ID, verify.item.Category.ID)
				assert.Equal(t, setup.item.Category.Name, verify.item.Category.Name)
				assert.Equal(t, setup.item.Tags[0].ID, verify.item.Tags[0].ID)
				assert.Equal(t, setup.item.Tags[0].Name, verify.item.Tags[0].Name)
				assert.Equal(t, setup.item.Tags[1].ID, verify.item.Tags[1].ID)
				assert.Equal(t, setup.item.Tags[1].Name, verify.item.Tags[1].Name)
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
