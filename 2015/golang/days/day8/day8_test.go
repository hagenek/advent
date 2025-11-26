package day8_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day8"
	"github.com/hagenek/advent/2015/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// TODO: Add test cases based on problem examples
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day8.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	t.Skip("Skipping until Part 1 is implemented")
	input, err := utils.ReadInput(8)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day8.Part1(input)
	assert.Greater(t, got, 0)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// TODO: Add test cases based on problem examples
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day8.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Skipping until Part 2 is implemented")
	input, err := utils.ReadInput(8)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day8.Part2(input)
	assert.Greater(t, got, 0)
}
