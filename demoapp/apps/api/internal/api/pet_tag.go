package api

type PetTag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AssertTagRequired checks if the required fields are not zero-ed
func AssertTagRequired(obj PetTag) error {
	return nil
}

// AssertTagConstraints checks if the values respects the defined constraints
func AssertTagConstraints(obj PetTag) error {
	return nil
}
