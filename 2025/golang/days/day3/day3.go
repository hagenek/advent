package day3

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func ParseLine(line string) int {

	max := 0

	for i := 0; i < len(line)-1; i++ {
		oi := string(line[i])
		for j := i + 1; j < len(line); j++ {
			ci := string(line[j])
			i1, err := strconv.Atoi(string(oi) + string(ci))
			if err != nil {
				fmt.Printf("Unable to convert %v \n to int", oi)
			}
			if i1 > max {
				max = i1
			}
		}
	}
	return max
}

func getFromFront(line []int) int {
	if len(line) <= 1 {
		return line[0]
	}
	if line[0] > line[1] {
		return line[0]
	} else {
		return -1
	}
}

func ParseLine2(line string) int {
	currLengthLine := len(line)

	fmt.Println("Current length:", currLengthLine)
	fmt.Println("Line:", line)
	iLine := []int{}

	for _, c := range line {
		i, _ := strconv.Atoi(string(c))
		iLine = append(iLine, i)
	}

	// how to capture value of number, priority to remove.
	// remove from front untill either len is 12 or number behind first is lower.

	optArr := []int{}

	for {

		if getFromFront(iLine) == -1 {
			optArr = append(optArr, iLine[0])
			iLine = iLine[1:]
		} else {
			iLine = slices.Concat(iLine[:1], iLine[2:])
			continue
		}

		if len(optArr)+len(iLine) == 12 {
			sum := ""
			for _, n := range optArr { // Note: fixed `range` - was `n := range`, should be `_, n := range`
				sum += strconv.Itoa(n)
			}
			summy, _ := strconv.Atoi(sum)
			return summy
		}
	}
}

// Part1 solves Advent of Code 2025 Day 3, Part 1.
// Implement the solution and return the result as an int.
func Part1(input string) int {
	total := 0
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		total += ParseLine(line)
	}
	return total
}

func Part2(input string) int {
	total := 0
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		total += ParseLine2(line)
	}
	return total
}
