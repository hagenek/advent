package day6_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day6"
	"github.com/hagenek/advent/2015/golang/utils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example test case",
			input: "example input",
			want:  0, // TODO: update with actual expected value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day6.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	input, err := utils.ReadInput(6)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day6.Part1(input)
	assert.Equal(t, 0, got) // TODO: update with actual expected value
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example test case",
			input: "example input",
			want:  0, // TODO: update with actual expected value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day6.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Day 6 Part 2 not implemented yet")
	input, err := utils.ReadInput(6)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day6.Part2(input)
	assert.Equal(t, 0, got) // TODO: update with actual expected value
}
