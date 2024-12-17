package solutions

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func makeSortedSlices(input string) ([]int, []int) {
	trimmed_input := strings.Trim(input, " \t\n\r")
	lines := strings.Split(trimmed_input, "\n")

	colA := make([]int, 0)
	colB := make([]int, 0)

	for i := range lines {
		a := strings.Split(lines[i], "   ")[0]
		b := strings.Split(lines[i], "   ")[1]
		ai, err := strconv.Atoi(a)
		if err != nil {
			continue
		}
		bi, err := strconv.Atoi(b)
		if err != nil {
			continue
		}

		colA = append(colA, ai)
		colB = append(colB, bi)
	}

	sort.Slice(colA, func(i, j int) bool {
		return colA[i] < colA[j]
	})
	sort.Slice(colB, func(i, j int) bool {
		return colB[i] < colB[j]
	})
	return colA, colB
}

func Day1Part1(input string) int {
	colA, colB := makeSortedSlices(input)

	count := 0
	for i := range colA {
		a := colA[i]
		b := colB[i]

		count += int(math.Abs(float64(b - a)))
	}
	return count
}

func Day1Part2(input string) int {
	colA, colB := makeSortedSlices(input)

	occurencesDict := make(map[int]int)

	for i := range colB {
		a, exists := occurencesDict[colB[i]]
		b := a + 1

		if exists {
			occurencesDict[colB[i]] = b
		} else {
			occurencesDict[colB[i]] = 1
		}
	}
	score := 0
	for i := range colA {
		a, exists := occurencesDict[colA[i]]
		if exists {
			score += a * colA[i]
		}
	}
	return score
}

func Day1() (int, int) {
	content, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		return 0, 0
	}
	inputString := string(content)

	return Day1Part1(inputString), Day1Part2(inputString)
}
