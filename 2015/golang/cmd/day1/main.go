package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day1"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	input, err := utils.ReadInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day1.Part1(input)
	fmt.Println(result)
	resultDayTwo := day1.Part2(input)
	fmt.Println("Day 2", resultDayTwo)
}
