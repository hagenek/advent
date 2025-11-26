package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day4"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(4)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day4.Part1(inp)
	result2 := day4.Part2(inp)
	fmt.Printf("Day1: %d\n", result)
	fmt.Printf("Day2: %d\n", result2)
}
