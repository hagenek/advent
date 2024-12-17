package solutions

import "testing"

func TestDay6Part1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	result := Day6Part1(input)
	expected := 41

	if result != expected {
		t.Errorf("Day6Part1(%v) = %d; want %d", input, result, expected)
	}
}
