package main

import (
	"fmt"

	"github.com/hagenek/advent/2025/golang/days/day1"
	"github.com/hagenek/advent/2025/golang/utils"
)

func main() {
	input, err := utils.ReadInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day1.Part1(input)
	fmt.Println("Part 1:", result)
	result2 := day1.Part2(input)
	fmt.Println("Part 2:", result2)
}