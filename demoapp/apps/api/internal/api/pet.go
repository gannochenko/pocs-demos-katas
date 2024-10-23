package api

import (
	"github.com/jinzhu/copier"

	"api/internal/domain"
)

type Pet struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	Category  Category `json:"category,omitempty"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags,omitempty"`
	Status    string   `json:"status,omitempty"`
}

func (p *Pet) ToDomain() (*domain.Pet, error) {
	result := &domain.Pet{}
	err := copier.Copy(result, p)
	if err != nil {
		return nil, err
	}

	return result, err
}
