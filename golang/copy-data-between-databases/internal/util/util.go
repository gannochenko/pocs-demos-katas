package util

import (
	"encoding/csv"
	"fmt"
	"os"

	"golang.org/x/term"
)

func PromptUser(prompt string) string {
	fmt.Print(prompt + ": ")
	input, _ := term.ReadPassword(0)

	return string(input)
}

func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
