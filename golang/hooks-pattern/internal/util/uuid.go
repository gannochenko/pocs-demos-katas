package util

import "github.com/google/uuid"

// GetUUID makes a new UUID v4
func GetUUID() string {
	return uuid.New().String()
}
