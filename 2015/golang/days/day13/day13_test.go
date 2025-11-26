package day13

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
			name: "Alice and Carol happiness",
			input: `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`,
			expected: 330, // TODO: Update expected value based on problem requirements
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
