package day8

import (
	"strconv"
	"strings"
)

// Part1 solves Advent of Code 2015 Day 8, Part 1.
func Part1(input string) int {
	codeLength := 0
	memoryLength := 0

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		codeLength += len(line)

		unquoted, err := strconv.Unquote(line)
		if err != nil {
			panic(err)
		}
		memoryLength += len(unquoted)
	}

	return codeLength - memoryLength
}

// Part2 solves Advent of Code 2015 Day 8, Part 2.
func Part2(input string) int {
	originalLength := 0
	encodedLength := 0

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		originalLength += len(line)

		// Quote adds surrounding quotes and escapes internal quotes and backslashes
		encoded := strconv.Quote(line)
		encodedLength += len(encoded)
	}

	return encodedLength - originalLength
}
