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
		t.Errorf("Day6Part1\n%v = %d; want %d", input, result, expected)
	}
}

func TestDay6Part1_2(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
#......#..`

	result := Day6Part1(input)
	expected := 43

	if result != expected {
		t.Errorf("Day6Part1\n%v = %d; want %d", input, result, expected)
	}
}

// func TestDay6Part2MissingSouthEast(t *testing.T) {
// 	input :=
// 		`0#000
// 1^11#
// 22222
// #3333
// 44444`

// 	result := Day6Part2(input)
// 	expected := 1

// 	if result != expected {
// 		t.Errorf("Day6Part2\n%v = %d; want %d", input, result, expected)
// 	}
// }

// func TestDay6Part2MissingNorthEast(t *testing.T) {
// 	input :=
// 		`0#000
// 1^111
// 22222
// #3333
// 444#4`

// 	result := Day6Part2(input)
// 	expected := 1

// 	if result != expected {
// 		t.Errorf("Day6Part2\n%v = %d; want %d", input, result, expected)
// 	}
// }

// func TestDay6Part2MissingSouthWest(t *testing.T) {
// 	input :=
// 		`0#000
// 1^11#
// 22222
// 33333
// 444#4`

// 	result := Day6Part2(input)
// 	expected := 1

// 	if result != expected {
// 		t.Errorf("Day6Part2\n%v = %d; want %d", input, result, expected)
// 	}
// }
