package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	// TODO: Add Part1 tests when requirements are known
	result := Part1("")
	assert.Equal(t, 0, result)
}

func TestIncreasesStraight(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"xyz", true},
		{"xyzabc", true},
		{"abcxyz", true},
		{"aabbcc", false},
		{"abd", false},
		{"ab", false},
		{"", false},
	}

	for _, tt := range tests {
		result := IncreasesStraight(tt.input)
		if result != tt.expected {
			t.Errorf("IncreasesStraight(%q) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

func TestNoConfusingLetters(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abc", true},
		{"xyz", true},
		{"abcdefghjkmnpqrstuvwxyz", true},
		{"", true},
		{"i", false},
		{"o", false},
		{"l", false},
		{"hello", false},
		{"lion", false},
		{"abci", false},
	}

	for _, tt := range tests {
		result := NoConfusingLetters(tt.input)
		if result != tt.expected {
			t.Errorf("NoConfusingLetters(%q) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

func TestHasTwoDistinctPairs(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aabb", true},
		{"zzabc", false},
		{"aabcc", true},
		{"aaa", false},
		{"aaaa", false},
		{"abcdee", false},
		{"aabaa", false},
		{"", false},
		{"aa", false},
		{"aabbcc", true},
	}

	for _, tt := range tests {
		result := AtLeastTwoDifferentPairs(tt.input)
		if result != tt.expected {
			t.Errorf("HasTwoDistinctPairs(%q) = %v, expected %v", tt.input, result, tt.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	// TODO: Add Part2 tests when requirements are known
	result := Part2("")
	assert.Equal(t, 0, result)
}
