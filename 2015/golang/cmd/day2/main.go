package main

import (
	"fmt"

	"github.com/hagenek/advent/2015/golang/days/day2"
	"github.com/hagenek/advent/2015/golang/utils"
)

func main() {
	input, err := utils.ReadInput(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := day2.Part1(input)
	fmt.Println(result)
}
