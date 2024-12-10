// days/day01_test.go
package days

import (
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	tests := []struct {
		name     string
		mass     int
		expected int
	}{
		{"mass of 12", 12, 2},
		{"mass of 14", 14, 2},
		{"mass of 1969", 1969, 654},
		{"mass of 100756", 100756, 33583},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateFuel(tt.mass)
			if got != tt.expected {
				t.Errorf("calculateFuel(%d) = %d; want %d", tt.mass, got, tt.expected)
			}
		})
	}
}

func TestDay1Part1(t *testing.T) {
	input := "12\n14\n1969\n100756"
	part1, _ := Day1(input)
	expected := 2 + 2 + 654 + 33583 // 34241

	if part1 != expected {
		t.Errorf("Day1 part 1 = %d; want %d", part1, expected)
	}
}
