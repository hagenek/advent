package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Short",
			input:    `R2, L3`,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Part1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		// TODO: Add test cases based on the problem examples
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Part2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
