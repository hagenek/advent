package days

import (
	"fmt"
	"os"
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
