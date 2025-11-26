package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day6"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	inp, err := utils.ReadInput(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := day6.Part1(inp)
	res2 := day6.Part2(inp)
	fmt.Printf("LIGHTS ON %d\n", res)
	fmt.Printf("LIGHTS ON PART 2: %d", res2)
}
