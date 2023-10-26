package day1

import (
	"strings"
)

// Part1 solves the first part of day 1 challenge.
func Part1(input string) int {
	floor := 0
	inputSplice := strings.Split(input, "")
	for _, value := range inputSplice {
		if value == "(" {
			floor++
		} else {
			floor--
		}
	}

	return floor
}
