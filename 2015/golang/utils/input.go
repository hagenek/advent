package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// findProjectRoot finds the project root by looking for go.mod
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", errors.New("could not find project root (go.mod not found)")
}

// ReadInput reads the input file for the specified day (1-24) and returns its content as a string.
// It returns an error if the day number is out of range, or if there's an error reading the file.
func ReadInput(day int) (string, error) {
	if day < 1 || day > 24 {
		return "", errors.New("day number must be between 1 and 24")
	}

	root, err := findProjectRoot()
	if err != nil {
		return "", fmt.Errorf("error finding project root: %v", err)
	}

	filePath := filepath.Join(root, "input", fmt.Sprintf("day%d.txt", day))
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	return string(data), nil
}

func SlidingWindow(s string, size int) []string {
	if size <= 0 || size > len(s) {
		return []string{}
	}
	windows := make([]string, 0, len(s)-size+1)
	for i := 0; i <= len(s)-size; i++ {
		windows = append(windows, s[i:i+size])
	}
	return windows
}

func SlidingWindowSlice[T any](slice []T, size int) [][]T {
	if size <= 0 || size > len(slice) {
		return [][]T{}
	}
	windows := make([][]T, 0, len(slice)-size+1)
	for i := 0; i <= len(slice)-size; i++ {
		window := make([]T, size)
		copy(window, slice[i:i+size])
		windows = append(windows, window)
	}
	return windows
}
