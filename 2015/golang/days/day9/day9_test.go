package day9_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day9"
	"github.com/hagenek/advent/2015/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1_Example(t *testing.T) {
	input := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`
	want := 605
	got := day9.Part1(input)
	assert.Equal(t, want, got)
}

func TestPart1_Input(t *testing.T) {
	t.Skip("Skipping until Part 1 is implemented")
	input, err := utils.ReadInput(9)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day9.Part1(input)
	assert.Greater(t, got, 0)
}

func TestPart2_Example(t *testing.T) {
	t.Skip("TODO: Add test case based on problem example")
	input := ""
	want := 0
	got := day9.Part2(input)
	assert.Equal(t, want, got)
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Skipping until Part 2 is implemented")
	input, err := utils.ReadInput(9)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day9.Part2(input)
	assert.Greater(t, got, 0)
}
