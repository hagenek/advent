package day5

import (
	"fmt"
	"strings"
)

// Part1 solves Advent of Code 2015 Day 5, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement

func stringIsNice(input string) bool {
	vowels := "aeiou"
	vowelCount := 0
	oneCharTwiceInARow := false
	var prevChar rune

	for idx, char := range input {

		if strings.ContainsRune(vowels, char) {
			vowelCount++
		}
		if idx > 0 && char == prevChar {
			oneCharTwiceInARow = true
		}
		charPair := string(prevChar) + string(char)
		switch charPair {
		case "ab", "cd", "pq", "xy":
			return false
		}
		prevChar = char
	}
	return vowelCount >= 3 && oneCharTwiceInARow
}

func Part1(input string) int {

	lines := strings.Split(input, "\n")
	niceLines := 0

	for _, line := range lines {
		if stringIsNice(line) {
			fmt.Printf("%s is nice", line)
			niceLines++
		}
	}
	return niceLines
}

func RemoveRepeating(inp string) string {
	var pc rune
	var ppc rune

	cleanedString := ""

	for idx, c := range inp {
		if idx < 2 {
			cleanedString = cleanedString + string(c)
			ppc = pc
			pc = c
			continue
		}
		if c == pc && c == ppc {
			continue
		}
		cleanedString = cleanedString + string(c)
		ppc = pc
		pc = c
	}
	return cleanedString
}

func hasRepeatWithOneBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func isNicePart2(charSeq string) bool {
	return hasPairTwice(charSeq) && hasRepeatWithOneBetween(charSeq)
}

func hasPairTwice(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pair := s[i:i+2]
		for j := i+2; j < len(s)-1; j++ {
			if s[j:j+2] == pair {
				return true
			}
		}
	}
	return false
}

// Part2 solves Advent of Code 2015 Day 5, Part 2.
// Implement the solution and return the result as an int.
// TODO: implement
func Part2(input string) int {
	fmt.Printf("Part2 called with input: %q\n", input)
	niceLineCount := 0
	for line := range strings.SplitSeq(input, "\n") {
		fmt.Printf("Processing line: %q\n", line)
		if isNicePart2(line) {
			fmt.Printf("Line %q is NICE\n", line)
			niceLineCount++
		} else {
			fmt.Printf("Line %q is NAUGHTY\n", line)
		}
	}
	fmt.Printf("Total nice lines: %d\n", niceLineCount)
	return niceLineCount
}
