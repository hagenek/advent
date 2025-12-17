package day17

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/utils"
)

// Part1 solves Advent of Code 2015 Day 17, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement
func Part1(input string) int {
	//
	change := []int{20, 15, 10, 5, 5}

	ps := utils.HeapsPermutation(change)

	for _, p := range ps {
		for _, c := range p {
			fmt.Printf("Value: %d", c)
		}
	}

	return 0
}

// Part2 solves Advent of Code 2015 Day 17, Part 2.
// Implement the solution and return the result as an int.
// TODO: implement
func Part2(input string) int {
	// TODO: implement part 2
	return 0
}
