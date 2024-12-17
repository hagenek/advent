package days

import (
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
			input: `ULL
RRDDD
LURDL
UUUUD`,
			expected: 1985,
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
		expected string
	}{
		{
			name: "Demo data part 2",
			input: `ULL
RRDDD
LURDL
UUUUD`,
			expected: "5DB3",
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
