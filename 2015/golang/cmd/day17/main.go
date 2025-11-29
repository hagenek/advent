package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day17"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(17)
	if err != nil {
		fmt.Println(err)
		return
	}
	result1 := day17.Part1(inp)
	result2 := day17.Part2(inp)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
