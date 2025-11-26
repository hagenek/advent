package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	// TODO: Update tests based on actual Part1 requirements
	result := Part1("1")
	assert.Equal(t, 2, result)
}

func TestExample2(t *testing.T) {
	result := Part1("11")
	assert.Equal(t, 2, result)
}

func TestExample3(t *testing.T) {
	result := Part1("21")
	assert.Equal(t, 4, result)
}

func TestExample4(t *testing.T) {
	result := ConvertStr("1211")
	assert.Equal(t, "111221", result)
}

func TestExample5(t *testing.T) {
	result := ConvertStr("111221")
	assert.Equal(t, "312211", result)
}

func TestPart2(t *testing.T) {
	// TODO: Add Part2 tests when requirements are known
}
