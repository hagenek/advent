package day12

import (
	"fmt"
	"regexp"
	"strconv"
)

// Part1 solves Advent of Code 2015 Day 12, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement
func Part1(input string) int {
	// TODO: implement part 1
	re := regexp.MustCompile(`-?\d+`)
	count := 0
	for _, m := range re.FindAllString(input, -1) {
		mi, _ := strconv.Atoi(m)
		count += mi
	}
	return count
}

// Part2 solves Advent of Code 2015 Day 12, Part 2.
// Implement the solution and return the result as an int.
func Part2(input string) int {
	// if we encounter "red" remove everything up untill } and backtowards and including }.
	// so lets store idx of last seen {, search forward for index of next } (by looking at substring of input starting at idx of prev })
	// then using those two indexes removing that part from the input, and run next search for red (function removeFirstRed)
	// removeRed will return the new string and a bool to say if it needs to continue or not
	// on string with red stuff removed, run part1

	for {
		newInput, shouldContinue := removeFirstRed(input)
		if !shouldContinue {
			break
		}
		input = newInput
	}

	return Part1(input)
}

// removeFirstRed removes the first occurrence of "red" along with its containing object
// returns the new string and a bool indicating if it needs to continue
func removeFirstRed(input string) (string, bool) {
	redIdx := -1
	for i := 0; i < len(input)-4; i++ {
		if input[i:i+5] == ":\"red" {
			redIdx = i
			break
		}
	}

	if redIdx == -1 {
		return input, false
	}

	// find last seen { before red
	startIdx := -1
	needsToCloseOpening := 0
	for i := redIdx - 1; i >= 0; i-- {
		if input[i] == '}' {
			needsToCloseOpening++
		}
		if input[i] == '{' {
			if needsToCloseOpening > 0 {
				needsToCloseOpening--
				continue
			}
			startIdx = i
			break
		}
	}

	// search forward for index of next } that doesnt have a {
	endIdx := -1
	needsToClose := 0
	for i := redIdx; i < len(input); i++ {
		if input[i] == '{' {
			needsToClose++
		}
		if input[i] == '}' {
			if needsToClose > 0 {
				needsToClose--
				continue
			}
			endIdx = i
			break
		}
	}

	// remove that part from the input
	fmt.Printf("Removing obj: %s\n", input[startIdx:endIdx+1])
	newInput := input[:startIdx] + input[endIdx+1:]
	fmt.Printf("New input: %s\n", newInput)
	return newInput, true
}
