package solutions

import (
	"os"
	"strings"
)

// Guard goes 1 and 1 ++ in direction untill next space has a #
// then turns 90 degrees right (changes direction) and starts going again
// while guard walks, add positions visited to a set (go map)? and then return len (size) of that map
// to solve the task

func Day6Part1(input string) int {
	lines := strings.Split(input, "\n")

	grid := LinesToRuneGrid(lines)

	return 0
}

func Day6Part2(input string) int {
	return 0
}

func Day6() (int, int) {
	content, err := os.ReadFile("inputs/day06.txt")
	if err != nil {
		return 0, 0
	}
	inputString := strings.TrimSpace(string(content))

	return Day6Part1(inputString), Day6Part2(inputString)
}
