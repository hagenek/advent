package days

import (
	"os"
	"strings"
)

func Day1Part1(input string) int {
	lines := strings.Split(input, "\n")

	for i := range lines {
		line := lines[i]
		println(line)
	}
	return 0
}

func Day1Part2(input string) int {
	return 0
}

func Day1() (int, int) {
	content, err := os.ReadFile("inputs/day02.txt")
	if err != nil {
		return 0, 0
	}
	inputString := string(content)

	return Day1Part1(inputString), Day1Part2(inputString)
}
