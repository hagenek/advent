package days

import (
	"strings"
)

func Day1Part1(input string) int {
	s := strings.Split(input, "")
	nums := Map(s, parseInt)
	ct := 0

	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		var next int
		if i+1 == len(nums) {
			next = nums[0]
		} else {
			next = nums[i+1]
		}

		if curr == next {
			ct += curr
		}
	}
	return ct
}

func Day1Part2(input string) int {
	s := strings.Split(input, "")
	nums := Map(s, parseInt)
	ct := 0

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		nextIndex := (i + len(nums)/2) % len(nums)
		next := nums[nextIndex]

		if curr == next {
			ct += curr
		}
	}
	return ct
}

func Day1() (int, int) {
	inputString := readInputFile(1)

	return Day1Part1(inputString), Day1Part2(inputString)
}
