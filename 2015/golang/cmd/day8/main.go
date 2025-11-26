package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day8"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(8)
	if err != nil {
		fmt.Println(err)
		return
	}
	result1 := day8.Part1(inp)
	result2 := day8.Part2(inp)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
