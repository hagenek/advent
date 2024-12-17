package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func listInputFiles() {
	entries, err := os.ReadDir("inputs")
	if err != nil {
		panic(fmt.Errorf("failed to read inputs directory: %w", err))
	}

	fmt.Println("\nFiles in inputs directory:")
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Printf("  - %s\n", entry.Name())
		}
	}
	fmt.Println()
}

func readInputFile(day int) string {
	data, err := os.ReadFile(fmt.Sprintf("inputs/day%d.txt", day))
	if err != nil {
		listInputFiles()
		panic(fmt.Errorf("failed to read input for day %d: %w", day, err))
	}
	return strings.TrimSpace(string(data))
}

func parseInt(r string) int {
	num, err := strconv.Atoi(r)
	if err != nil {
		fmt.Printf("Error converting %s, err: %s", r, err)
	}
	return num
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = fn(slice[i])
	}
	return result
}

func Reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
