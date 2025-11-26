package day5_test

import (
	"testing"

	"github.com/hagenek/advent/2015/golang/days/day5"
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
			name:  "nice string - ugknbfddgicrmopn",
			input: "ugknbfddgicrmopn",
			want:  1,
		},
		{
			name:  "nice string - aaa",
			input: "aaa",
			want:  1,
		},
		{
			name:  "naughty string - jchzalrnumimnmhp",
			input: "jchzalrnumimnmhp",
			want:  0,
		},
		{
			name:  "naughty string - haegwjzuvuyypxyu",
			input: "haegwjzuvuyypxyu",
			want:  0,
		},
		{
			name:  "naughty string - dvszwmarrgswjxmb",
			input: "dvszwmarrgswjxmb",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.Part1(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart1_Input(t *testing.T) {
	input, err := utils.ReadInput(5)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day5.Part1(input)
	assert.Equal(t, 255, got)
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "nice string - xxyxx (has xx twice and xyx)",
			input: "xxyxx",
			want:  1,
		},

		{
			name:  "naughty - has triple aaa but no repeating pair)",
			input: "csuellnlcyqaaamq",
			want:  0,
		},
		{
			name:  "nice string - xxyxx (has xx twice and xyx)",
			input: "xxyxx",
			want:  1,
		},
		{
			name:  "naughty string - xxxdokaf triple but not double",
			input: "xxxydf3rx",
			want:  0,
		},
		{
			name:  "naughty string - xxxdokaf triple but not double",
			input: "xeqerAAAdfsdo",
			want:  0,
		},
		{
			name:  "nice - has pair twice AND xyx pattern",
			input: "aabaa",
			want:  1,
			// "aa" appears at 0-1 and 3-4 (non-overlapping)
			// and "aba" pattern exists at 1-2-3
		},
		{
			name:  "naughty string - ieodomkazucvgmuy (has odo but no pair twice)",
			input: "ieodomkazucvgmuy",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.Part2(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRemoveRepeating(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no repeating chars",
			input: "abcdef",
			want:  "abcdef",
		},
		{
			name:  "triple a at start",
			input: "aaabcd",
			want:  "aabcd",
		},
		{
			name:  "triple x in middle",
			input: "abxxxef",
			want:  "abxxef",
		},
		{
			name:  "triple z at end",
			input: "abcdezzz",
			want:  "abcdezz",
		},
		{
			name:  "multiple triples",
			input: "aaabbbbcccc",
			want:  "aabbcc",
		},
		{
			name:  "quad becomes double",
			input: "xxxxyz",
			want:  "xxyz",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "single char",
			input: "a",
			want:  "a",
		},
		{
			name:  "double char",
			input: "aa",
			want:  "aa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.RemoveRepeating(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPart2_Input(t *testing.T) {
	t.Skip("Day 5 Part 2 not implemented yet")
	input, err := utils.ReadInput(5)
	if err != nil {
		t.Fatalf("failed reading input: %v", err)
	}
	got := day5.Part2(input)
	assert.Equal(t, 0, got)
}
