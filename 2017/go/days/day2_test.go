package days

import (
	"sort"
	"strings"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example_1",
			input: `5 1 9 5
7 5 3
2 4 6 8`,
			expected: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Day2Part1(tt.input)
			if result != tt.expected {
				t.Errorf("Day2Part1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example_1",
			input: `5 9 2 8
9 4 7 3
3 8 6 5`,
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Day2Part2(tt.input)
			if result != tt.expected {
				t.Errorf("Day2Part2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Original version
func Day2Part2Original(input string) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		ns := strings.Fields(line)
		nums := Map(ns, parseInt)
		for i, x := range nums {
			for _, y := range nums[i+1:] {
				if x%y == 0 {
					total += x / y
				} else if y%x == 0 {
					total += y / x
				}
			}
		}
	}
	return total
}

// Sorted version
func Day2Part2Sorted(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		ns := strings.Fields(line)
		nums := Map(ns, parseInt)
		sort.Sort(sort.Reverse(sort.IntSlice(nums)))

		for i, x := range nums {
			for _, y := range nums[i+1:] {
				if x%y == 0 {
					total += x / y
					break
				}
			}
		}
	}
	return total
}

// Test data generator
func generateTestData(size int) string {
	var sb strings.Builder
	for i := 0; i < size; i++ {
		// Create rows with 8 numbers each
		sb.WriteString("16 1024 2048 4096 8192 16384 32768 65536\n")
	}
	return sb.String()
}

func BenchmarkDay2Part2(b *testing.B) {
	// Generate test data with 1000 rows
	input := generateTestData(1000)

	b.Run("Original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Day2Part2Original(input)
		}
	})

	b.Run("Sorted", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Day2Part2Sorted(input)
		}
	})
}
