package api

type Customer struct {
	Id       int64     `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Address  []Address `json:"address,omitempty"`
}

// AssertCustomerRequired checks if the required fields are not zero-ed
func AssertCustomerRequired(obj Customer) error {
	for _, el := range obj.Address {
		if err := AssertAddressRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCustomerConstraints checks if the values respects the defined constraints
func AssertCustomerConstraints(obj Customer) error {
	for _, el := range obj.Address {
		if err := AssertAddressConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
