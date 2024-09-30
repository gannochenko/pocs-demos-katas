package api

type Category struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AssertCategoryRequired checks if the required fields are not zero-ed
func AssertCategoryRequired(obj Category) error {
	return nil
}

// AssertCategoryConstraints checks if the values respects the defined constraints
func AssertCategoryConstraints(obj Category) error {
	return nil
}
