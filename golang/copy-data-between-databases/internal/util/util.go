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
	fmt.Print("\n")

	return string(input)
}

func ReadCSV(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []string

	for _, row := range rows {
		records = append(records, row[0])
	}

	return records, nil
}
