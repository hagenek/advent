// days/day01_test.go
package days

import (
	"testing"
)

// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
func TestDay2Part1(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		expected []int
	}{
		{"example 1", "1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"example 2", "2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"example 3", "2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"example 4", "1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"original test", "1,9,10,3,2,3,11,0,99,30,40,50", []int{3500, 9, 10, 70,
			2, 3, 11, 0,
			99,
			30, 40, 50}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Day2Solutions.Part1(tt.program)
			if tt.expected[0] != got {
				t.Errorf("Day2Part1(%s) = %v; \n want %v", tt.program, got, tt.expected[0])
			}
		})
	}
}
