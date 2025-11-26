package main

import (
	"fmt"
	"strings"

	"github.com/hagenek/advent/2015/golang/days/day10"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(10)
	if err != nil {
		fmt.Println(err)
		return
	}
	inpT := strings.TrimSpace(inp)
	result1 := day10.Part1(inpT)
	result2 := day10.Part2(inpT)
	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part2: %d\n", result2)
}
