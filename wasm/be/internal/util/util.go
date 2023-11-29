package util

import (
	"be/internal/constants"
)

func GetSampleConstant() string {
	return constants.SampleConstant
}

type SamplePerson struct {
	Name   string
	IsEvil bool
	Habits map[string]string
}

func GetPerson() *SamplePerson {
	return &SamplePerson{
		Name:   "Baggins",
		IsEvil: false,
		Habits: map[string]string{
			"catch_phrase": "where is my ring?",
		},
	}
}
