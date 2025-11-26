package main

import (
	"fmt"
	"strings"

	"github.com/hagenek/advent/2015/golang/days/day11"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(11)
	if err != nil {
		fmt.Println(err)
		return
	}
	inpT := strings.TrimSpace(inp)
	result1 := day11.Part1(inpT)
	result2 := day11.Part2(inpT)
	fmt.Printf("Part1: %s\n", result1)
	fmt.Printf("Part2: %s\n", result2)
}
