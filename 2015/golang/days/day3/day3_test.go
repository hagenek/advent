package day3_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day3"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single move",
			input: ">",
			want:  2,
		},
		{
			name:  "multiple moves with revisits",
			input: ">>>^v^",
			want:  5,
		},
		{
			name:  "back and forth",
			input: "^v^v^v^v^v",
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single move",
			input: ">",
			want:  2,
		},
		{
			name:  "Multiple moves",
			input: "^v",
			want:  3,
		},
		{
			name:  "back and forth",
			input: "^v^v^v^v^v",
			want:  11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
