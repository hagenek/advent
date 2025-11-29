package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day14"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	input, err := utils.ReadInput(14)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day14.Part1(input)
	fmt.Println(result)
	resultPart2 := day14.Part2(input)
	fmt.Println("Part 2:", resultPart2)
}