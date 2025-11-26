package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	dayTemplate = `package day{{DAY}}

// Part1 solves Advent of Code 2015 Day {{DAY}}, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement
func Part1(input string) int {
	// TODO: implement part 1
	return 0
}

// Part2 solves Advent of Code 2015 Day {{DAY}}, Part 2.
// Implement the solution and return the result as an int.
// TODO: implement
func Part2(input string) int {
	// TODO: implement part 2
	return 0
}
`

	testTemplate = `package day{{DAY}}_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day{{DAY}}"
	"github.com/hagenek/advent/2015/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example 1",
			input: "",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day{{DAY}}.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	t.Skip("Day {{DAY}} Part 1 not implemented yet")
	input, err := utils.ReadInput({{DAY}})
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day{{DAY}}.Part1(input)
	assert.Equal(t, 0, got)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example 1",
			input: "",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day{{DAY}}.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Day {{DAY}} Part 2 not implemented yet")
	input, err := utils.ReadInput({{DAY}})
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day{{DAY}}.Part2(input)
	assert.Equal(t, 0, got)
}
`

	mainTemplate = `package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day{{DAY}}"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput({{DAY}})
	if err != nil {
		fmt.Println(err)
		return
	}
	result1 := day{{DAY}}.Part1(inp)
	result2 := day{{DAY}}.Part2(inp)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
`
)

func createDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	return nil
}

func writeFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}

func generateScaffolding(day int) error {
	dayStr := strconv.Itoa(day)

	// Create directories
	dayDir := filepath.Join("days", fmt.Sprintf("day%d", day))
	cmdDir := filepath.Join("cmd", fmt.Sprintf("day%d", day))

	if err := createDir(dayDir); err != nil {
		return err
	}
	if err := createDir(cmdDir); err != nil {
		return err
	}
	if err := createDir("input"); err != nil {
		return err
	}

	// Generate file contents by replacing template variables
	dayContent := strings.ReplaceAll(dayTemplate, "{{DAY}}", dayStr)
	testContent := strings.ReplaceAll(testTemplate, "{{DAY}}", dayStr)
	mainContent := strings.ReplaceAll(mainTemplate, "{{DAY}}", dayStr)

	// Write solution file
	dayFile := filepath.Join(dayDir, fmt.Sprintf("day%d.go", day))
	if err := writeFile(dayFile, dayContent); err != nil {
		return err
	}

	// Write test file
	testFile := filepath.Join(dayDir, fmt.Sprintf("day%d_test.go", day))
	if err := writeFile(testFile, testContent); err != nil {
		return err
	}

	// Write main file
	mainFile := filepath.Join(cmdDir, "main.go")
	if err := writeFile(mainFile, mainContent); err != nil {
		return err
	}

	// Create empty input file
	inputFile := filepath.Join("input", fmt.Sprintf("day%d.txt", day))
	if err := writeFile(inputFile, ""); err != nil {
		return err
	}

	fmt.Printf("Generated scaffolding for day %d:\n", day)
	fmt.Printf("  - %s\n", dayFile)
	fmt.Printf("  - %s\n", testFile)
	fmt.Printf("  - %s\n", mainFile)
	fmt.Printf("  - %s\n", inputFile)

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <day>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s 12\n", os.Args[0])
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Fprintf(os.Stderr, "Error: day must be a number between 1 and 25\n")
		os.Exit(1)
	}

	if err := generateScaffolding(day); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nTo get started:\n")
	fmt.Printf("  1. Add your puzzle input to input/day%d.txt\n", day)
	fmt.Printf("  2. Implement solutions in days/day%d/day%d.go\n", day, day)
	fmt.Printf("  3. Add test cases to days/day%d/day%d_test.go\n", day, day)
	fmt.Printf("  4. Run tests: go test ./days/day%d\n", day)
	fmt.Printf("  5. Run solution: go run cmd/day%d/main.go\n", day)
}