package util

func FromPtr[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}

	return defaultValue
}

// ToPtr allows creating a pointer value from a constant in-place
func ToPtr[T any](value T) *T {
	return &value
}
