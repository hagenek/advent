package day12_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day12"
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
			input: `{"e":"violet","a":-44,"d":115,","c"`,
			want:  71,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day12.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	t.Skip("Day 12 Part 1 not implemented yet")
	input, err := utils.ReadInput(12)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day12.Part1(input)
	assert.Equal(t, 0, got)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Test with small json",
			input: `[["green",[{"e":"green","a":1000,"d":{"c":"violet","a":"yellow","b":"violet"},"c":"yellow","b":1,"g":{"a":["yellow",-1],"r": {"e":5999, "e":"red"}}`,
			want:  1000,
		},
		{
			name:  "red in outer object with inner object before it",
			input: `{"a":{"b":5},"c":"red","d":10}`,
			want:  0,
		},
		{
			name:  "nested red should only remove inner",
			input: `{"a":100,"b":{"c":"red","d":5}}`,
			want:  100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day12.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Day 12 Part 2 not implemented yet")
	input, err := utils.ReadInput(12)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day12.Part2(input)
	assert.Equal(t, 0, got)
}
