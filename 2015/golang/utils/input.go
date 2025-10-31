package utils

import (
	"errors"
	"fmt"
	"os"
)

// ReadInput reads the input file for the specified day (1-24) and returns its content as a string.
// It returns an error if the day number is out of range, or if there's an error reading the file.
func ReadInput(day int) (string, error) {
	if day < 1 || day > 24 {
		return "", errors.New("day number must be between 1 and 24")
	}

	filepath := fmt.Sprintf("input/day%d.txt", day)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v. \nPlease ensure the file exists in the expected folder", err)
	}

	return string(data), nil
}
