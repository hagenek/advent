package day17_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day17"
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
			name:  "example 1",
			input: "",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day17.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	t.Skip("Day 17 Part 1 not implemented yet")
	input, err := utils.ReadInput(17)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day17.Part1(input)
	assert.Equal(t, 0, got)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example 1",
			input: "",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day17.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Day 17 Part 2 not implemented yet")
	input, err := utils.ReadInput(17)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day17.Part2(input)
	assert.Equal(t, 0, got)
}
