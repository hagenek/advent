package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day13"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	input, err := utils.ReadInput(13)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day13.Part1(input)
	fmt.Println(result)
	resultPart2 := day13.Part2(input)
	fmt.Println("Part 2:", resultPart2)
}