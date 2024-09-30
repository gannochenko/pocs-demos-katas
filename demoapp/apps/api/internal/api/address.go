package api

type Address struct {
	Street string `json:"street,omitempty"`
	City   string `json:"city,omitempty"`
	State  string `json:"state,omitempty"`
	Zip    string `json:"zip,omitempty"`
}

// AssertAddressRequired checks if the required fields are not zero-ed
func AssertAddressRequired(obj Address) error {
	return nil
}

// AssertAddressConstraints checks if the values respects the defined constraints
func AssertAddressConstraints(obj Address) error {
	return nil
}
