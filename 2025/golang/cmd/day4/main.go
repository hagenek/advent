package main

import (
	"fmt"

	"github.com/hagenek/advent/2025/golang/days/day4"
	"github.com/hagenek/advent/2025/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(4)
	if err != nil {
		fmt.Println(err)
		return
	}
	result1 := day4.Part1(inp)
	result2 := day4.Part2(inp)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
