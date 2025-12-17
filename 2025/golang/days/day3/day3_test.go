package day3_test

import (
	"testing"

	"github.com/hagenek/advent/2025/golang/days/day3"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"test1", "987654321111111", 98},
		{"test2", "811111111111119", 89},
		{"test3", "234234234234278", 78},
		{"test4", "818181911112111", 92},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.ParseLine(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestParseLine2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"test1", "987654321111111", 987654321111},
		{"test2", "811111111111119", 811111111119},
		{"test3", "234234234234278", 434234234278},
		{"test4", "218181911112111", 888911112111},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.ParseLine2(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}
