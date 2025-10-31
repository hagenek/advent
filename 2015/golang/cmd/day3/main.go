package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/day3"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day3.Part1(inp)
	fmt.Println(result)
	resultDayTwo := day3.Part2(inp)
	fmt.Println("Day 2", resultDayTwo)
}
