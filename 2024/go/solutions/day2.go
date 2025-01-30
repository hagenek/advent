package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

func diff(a int, b int) int {
	return int(math.Abs(float64(b - a)))
}

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")

	numberRows := Map(lines, func(line string) []int {
		nums := strings.Fields(line)
		return Map(nums, func(n string) int {
			num, _ := strconv.Atoi(n)
			return num
		})
	})

	okLines := 0

	// ascending check
	va := Map(numberRows, func(row []int) bool {
		for i := 0; i < len(row)-1; i++ {
			if !(row[i+1] > row[i] && diff(row[i+1], row[i]) < 4) {
				return false
			}
		}
		return true
	})

	for _, b := range va {
		if b {
			okLines++
		}
	}

	// descending check
	vd := Map(numberRows, func(row []int) bool {
		for i := 0; i < len(row)-1; i++ {
			if !(row[i+1] < row[i] && diff(row[i+1], row[i]) < 4) {
				return false
			}
		}
		return true
	})

	for _, b := range vd {
		if b {
			okLines++
		}
	}

	return okLines

}

func Day2Part2(input string) int {
	lines := strings.Split(input, "\n")

	numbers := Map(lines, func(line string) []int {
		nums := strings.Fields(line)
		return Map(nums, func(n string) int {
			num, _ := strconv.Atoi(n)
			return num
		})
	})

	okLines := 0

	for _, nums := range numbers {
		isFirst := true
		isAscending := false
		noMistake := true
		for i := 0; i < len(nums)-1; i++ {

			isLast := i == len(nums)-2
			next := nums[i+1]
			now := nums[i]
			diff := next - now

			// diff negative
			if diff < 0 && isFirst {
				isFirst = false
			}
			// diff positive
			if diff > 0 && isFirst {
				isAscending = true
				isFirst = false
			}

			if diff == 0 || diff > 4 || diff < -3 ||
				isAscending && diff < 0 ||
				!isAscending && diff > 0 && noMistake {
				noMistake = false

				if isLast {
					now = nums[i-1]
					diff = next - now
				} else {
					next = nums[i+2]
					diff = next - now
					i++
				}
			}

			if isAscending && diff < 0 {
				break
			}

			if !isAscending && diff > 0 {
				break
			}

			if diff == 0 || diff > 4 || diff < -3 {
				break
			}
			// diff negative
			if diff < 0 && isFirst {
				isAscending = false
				isFirst = false
			}
			// diff positive
			if diff > 0 && isFirst {
				isAscending = true
				isFirst = false
			}

			if isLast {
				fmt.Printf("Safe line: %v\n", nums)
				okLines++
			}
		}
	}
	return okLines
}

func Day2() (int, int) {

	inputString := readInputFile(2)

	return Day2Part1(inputString), Day2Part2(inputString)
}
