package days

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example_1",
			input:    "+1\n-2\n+3\n+1",
			expected: 3,
		},
		{
			name:     "example_2",
			input:    "+1\n+1\n+1",
			expected: 3,
		},
		{
			name:     "example_3",
			input:    "+1\n+1\n-2",
			expected: 0,
		},
		{
			name:     "example_4",
			input:    "-1\n-2\n-3",
			expected: -6,
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
		// {
		// 	name:     "example_2",
		// 	input:    "+1\n-1",
		// 	expected: 0,
		// },
		{
			name:     "example_3",
			input:    "+3\n+3\n+4s\n-2\n-4",
			expected: 10,
		},
		// {
		// 	name:     "example_4",
		// 	input:    "-6\n+3\n+8\n+5\n-6",
		// 	expected: 5,
		// },
		// {
		// 	name:     "example_5",
		// 	input:    "+7\n+7\n-2\n-7\n-4",
		// 	expected: 14,
		// },
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
