package main

import (
	"aoc2019-golang/days"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	day := flag.Int("day", 1, "day of advent to run")
	flag.Parse()

	inputFile := filepath.Join("inputs", fmt.Sprintf("day%02d.txt", *day))

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	cleanInput := strings.TrimSpace(string(input))

	switch *day {
	case 1:
		result1, result2 := days.Day1(cleanInput)
		fmt.Printf("Day %d Part 1: %v\n", *day, result1)
		fmt.Printf("Day %d Part 2: %v\n", *day, result2)
	default:
		fmt.Printf("Day %d not implemented\n", *day)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
