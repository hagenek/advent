package main

import (
	"testing"
)

func TestCountUniqueHouses(t *testing.T) {
	testCases := []struct {
		directions string
		expected   int
		desc       string
	}{
		{">", 2, "delivers presents to 2 houses: one at starting location, one to the east"},
		{"^>v<", 4, "delivers presents to 4 houses in a square"},
		{"^v^v^v^v^v", 2, "delivers presents to 2 houses repeatedly"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := countUniqueHouses(tc.directions)
			if result != tc.expected {
				t.Errorf("For directions %q, expected %d houses, got %d", tc.directions, tc.expected, result)
			}
		})
	}
}