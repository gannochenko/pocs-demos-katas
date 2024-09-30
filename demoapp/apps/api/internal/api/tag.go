package api

type Tag struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AssertTagRequired checks if the required fields are not zero-ed
func AssertTagRequired(obj Tag) error {
	return nil
}

// AssertTagConstraints checks if the values respects the defined constraints
func AssertTagConstraints(obj Tag) error {
	return nil
}
