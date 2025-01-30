package days

import (
	"math"
	"slices"
	"strings"
)

// split on lines
// split on fields
// Map to ints
// sort the lines
// Map math.Abs(max - min)
// Sum

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")
	totalDiff := 0.0

	for _, line := range lines {
		ns := strings.Fields(line)
		nums := Map(ns, parseInt)

		max := slices.Max(nums)
		min := slices.Min(nums)
		totalDiff += math.Abs(float64(max - min))
	}
	return int(totalDiff)
}
func Day2Part2(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		ns := strings.Fields(line)
		nums := Map(ns, parseInt)

		for i, x := range nums {
			for _, y := range nums[i+1:] {
				if x%y == 0 {
					total += x / y
				} else if y%x == 0 {
					total += y / x
				}
			}
		}
	}
	return total
}

func Day2() (int, int) {
	inputString := readInputFile(2)

	return Day2Part1(inputString), Day2Part2(inputString)
}
