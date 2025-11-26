package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day7"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	input, _ := utils.ReadInput(7)

	part1Result := day7.Part1(input)
	fmt.Printf("Part 1: %d\n", part1Result)

	part2Result := day7.Part2(input)
	fmt.Printf("Part 2: %d\n", part2Result)
}
