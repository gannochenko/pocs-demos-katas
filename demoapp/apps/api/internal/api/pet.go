package api

import (
	"github.com/jinzhu/copier"

	"api/internal/domain"
	httpUtil "api/internal/util/http"
	"api/pkg/syserr"
)

type Pet struct {
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name"`
	Category  PetCategory `json:"category,omitempty"`
	PhotoUrls []string    `json:"photoUrls"`
	Tags      []PetTag    `json:"tags,omitempty"`
	Status    string      `json:"status,omitempty"`
}

// AssertPetRequired checks if the required fields are not zero-ed
func AssertPetRequired(obj Pet) error {
	elements := map[string]interface{}{
		"name":      obj.Name,
		"photoUrls": obj.PhotoUrls,
	}
	for name, el := range elements {
		if isZero := httpUtil.IsZeroValue(el); isZero {
			return syserr.NewBadInput("required field missing", syserr.F("field", name))
		}
	}

	if err := AssertCategoryRequired(obj.Category); err != nil {
		return err
	}
	for _, el := range obj.Tags {
		if err := AssertTagRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPetConstraints checks if the values respects the defined constraints
func AssertPetConstraints(obj Pet) error {
	if err := AssertCategoryConstraints(obj.Category); err != nil {
		return err
	}
	for _, el := range obj.Tags {
		if err := AssertTagConstraints(el); err != nil {
			return err
		}
	}
	return nil
}

func (p *Pet) ToDomain() (*domain.Pet, error) {
	result := &domain.Pet{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, err
}
