package day11

import (
	"fmt"
	"strings"
	"time"
)

func IncreasesStraight(inp string) bool {
	// valid triplet is a rune that
	// b - a == 1 and c - b == 1

	for i := 0; i < len(inp)-2; i++ {
		if inp[i+1]-inp[i] == 1 && inp[i+2]-inp[i+1] == 1 {
			return true
		}
	}
	return false
}

func NoConfusingLetters(inp string) bool {
	// i, o, l not allowed

	for l := range inp {
		if strings.Contains("iol", string(inp[l])) {
			return false
		}
	}
	return true
}

func AtLeastTwoDifferentPairs(inp string) bool {
	pairCount := 0
	seenPairs := map[string]bool{}
	for i := 0; i < len(inp)-1; i++ {
		pair := string(inp[i : i+2])
		if inp[i] == inp[i+1] && !seenPairs[pair] {
			seenPairs[pair] = true
			pairCount++
		}
	}
	return pairCount >= 2
}

func IncrementPassword(inp string) string {
	// just find out rune value of a and z and check if letter is z, if it is
	// increment to a int32 value, or else increment by 1. if you go from z to a
	// try to increment next letter in string, untill a letter meets the increment by one
	// logic
	for i := len(inp) - 1; i >= 0; i-- {
		if inp[i] == 'z' {
			inp = inp[:i] + "a" + inp[i+1:]
		} else {
			inp = inp[:i] + string(inp[i]+1) + inp[i+1:]
			break
		}
	}
	return inp
}

// Part1 solves Advent of Code 2015 Day 11, Part 1.
func Part1(input string) string {
	// TODO: Implement Part1
	startTime := time.Now()
	pw := input
	for {
		if time.Since(startTime) > time.Second {
			panic("Couldnt find valid password after 1 second, logic fault")
		}
		pw = IncrementPassword(strings.TrimSpace(pw))
		fmt.Printf("Pw: %s\n", pw)
		if IncreasesStraight(pw) && AtLeastTwoDifferentPairs(pw) && NoConfusingLetters(pw) {
			return pw
		}
	}
}

// Part2 solves Advent of Code 2015 Day 11, Part 2.
func Part2(input string) string {
	// TODO: Implement Part2
	return Part1(Part1(input))
}
