package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	testInput := `test input here`
	expected := 0 // Update with expected result
	result := Part1(testInput)
	assert.Equal(t, expected, result)
}

func TestPart1Example(t *testing.T) {
	testInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := 3
	result := Part1(testInput)
	assert.Equal(t, expected, result)
}

func TestPart2(t *testing.T) {
	testInput := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	expected := 6
	result := Part2(testInput)
	assert.Equal(t, expected, result)
}
