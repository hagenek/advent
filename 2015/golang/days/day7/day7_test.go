package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name: "example 1",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			expected: 0, // TODO: update with correct expected value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part1(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "example 1",
			input:    "test input", // TODO: add real test input
			expected: 0,            // TODO: update with correct expected value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Part2(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}