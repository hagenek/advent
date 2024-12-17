package days

import (
	"testing"
)

// For example:

// Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
// R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
// R5, L5, R5, R3 leaves you 12 blocks away.

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example_1",
			input:    `R2, L3`,
			expected: 5,
		},
		{
			name:     "example_2",
			input:    `R2, R2, R2`,
			expected: 2,
		},
		{
			name:     "example_3",
			input:    `R5, L5, R5, R3`,
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Day1Part1(tt.input)
			if result != tt.expected {
				t.Errorf("Day1Part1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDay1Part2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example_1",
			input:    `R8, R4, R4, R8`,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Day1Part2(tt.input)
			if result != tt.expected {
				t.Errorf("Day1Part2() = %v, want %v", result, tt.expected)
			}
		})
	}
}
