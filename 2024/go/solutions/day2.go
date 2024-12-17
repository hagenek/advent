package solutions

import (
	"os"
	"strings"
)

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")

	for i := range lines {
		line := lines[i]
		println(line)
	}
	return 0
}

func Day2Part2(input string) int {
	return 0
}

func Day2() (int, int) {
	content, err := os.ReadFile("inputs/day02.txt")
	if err != nil {
		return 0, 0
	}
	inputString := string(content)

	return Day2Part1(inputString), Day2Part2(inputString)
}
