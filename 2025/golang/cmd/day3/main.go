package main

import (
	"fmt"

	"github.com/hagenek/advent/2025/golang/days/day3"
	"github.com/hagenek/advent/2025/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	result1 := day3.Part1(inp)
	result2 := day3.Part2(inp)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
