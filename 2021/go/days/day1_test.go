package days

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
 tests := []struct {
  name     string
  input    string
  expected int
 }{
  {
   name:     "example_1",
   input:    `test input`,
   expected: 0,
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   result := Day1Part1(tt.input)
   if result != tt.expected {
    t.Errorf("Day1Part1() = %v, want %v", result, tt.expected)
   }
  })
 }
}

func TestDay1Part2(t *testing.T) {
 tests := []struct {
  name     string
  input    string
  expected int
 }{
  {
   name:     "example_1",
   input:    `test input`,
   expected: 0,
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   result := Day1Part2(tt.input)
   if result != tt.expected {
    t.Errorf("Day1Part2() = %v, want %v", result, tt.expected)
   }
  })
 }
}
