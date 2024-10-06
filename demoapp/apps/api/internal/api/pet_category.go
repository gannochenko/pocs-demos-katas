package api

type PetCategory struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AssertCategoryRequired checks if the required fields are not zero-ed
func AssertCategoryRequired(obj PetCategory) error {
	return nil
}

// AssertCategoryConstraints checks if the values respects the defined constraints
func AssertCategoryConstraints(obj PetCategory) error {
	return nil
}
