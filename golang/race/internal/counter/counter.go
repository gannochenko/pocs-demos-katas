package counter

var count int

// Increment increases the count
func Increment() {
	count++ // Data race occurs here
}

// GetCount returns the current count
func GetCount() int {
	return count
}
