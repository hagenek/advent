package day2_test

import (
	"github.com/hagenek/advent/2015/golang/pkg/day2"
	"github.com/hagenek/advent/2015/golang/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	input, err := utils.ReadInput(2)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day2.Part1(input)
	assert.Equal(t, 1586300, got)
}

func TestPart2(t *testing.T) {
	input, err := utils.ReadInput(2)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day2.Part2(input)
	assert.Equal(t, 3737498, got)
}
