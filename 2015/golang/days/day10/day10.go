package day10

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ConvertStr(input string) string {
	var sb strings.Builder
	sb.Grow(7000000)
	currCount := 0
	var prev rune
	for i, c := range input {
		if i == 0 && i != len(input)-1 {
			prev = c
			currCount = 1
			continue
		}
		if prev == c && i != len(input)-1 {
			currCount++
			prev = c
			continue
		} else if prev == c && i == len(input)-1 {
			currCount++
			sb.WriteString(strconv.Itoa(currCount))
			sb.WriteRune(prev)
		} else {
			sb.WriteString(strconv.Itoa(currCount))
			sb.WriteRune(prev)
			currCount = 1
			prev = c
			if i == len(input)-1 {
				sb.WriteString(strconv.Itoa(currCount))
				sb.WriteRune(prev)
			}
		}
	}
	return sb.String()
}

// Part1 solves Advent of Code 2015 Day 10, Part 1.
func Part1(input string) int {
	// TODO: Implement Part1
	// extract pairs

	for i := range 40 {
		iterStart := time.Now()
		input = ConvertStr(input)
		fmt.Printf("Iteration %d: %v (len: %d)\n", i+1, time.Since(iterStart), len(input))
	}
	return len(input)
}

// Part2 solves Advent of Code 2015 Day 10, Part 2.
func Part2(input string) int {
	start := time.Now()
	for i := range 50 {
		iterStart := time.Now()
		input = ConvertStr(input)
		fmt.Printf("Iteration %d: %v (len: %d)\n", i+1, time.Since(iterStart), len(input))
	}
	fmt.Println("Total:", time.Since(start))
	return len(input)
}
