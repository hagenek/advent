package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func handleError(err1 error, err2 error, err3 error) {
	if err1 != nil {
		fmt.Println("Error converting l:", err1)
		return
	}

	if err2 != nil {
		fmt.Println("Error converting w:", err2)
		return
	}

	if err3 != nil {
		fmt.Println("Error converting h:", err3)
		return
	}
}

func parseInput(input string) ([][]int, error) {
	inputSliced := strings.Split(input, "\n")
	var intAreaSlice [][]int

	for _, area := range inputSliced {
		if area == "" {
			continue
		}
		sliced := strings.Split(area, "x")

		if len(sliced) < 3 {
			return nil, fmt.Errorf("Not enough dimensions")
		}

		l, err1 := strconv.Atoi(sliced[0])
		w, err2 := strconv.Atoi(sliced[1])
		h, err3 := strconv.Atoi(sliced[2])

		if err1 != nil || err2 != nil || err3 != nil {
			return nil, fmt.Errorf("Error converting dimensions to integer")
		}

		intAreaSlice = append(intAreaSlice, []int{l, w, h})
	}

	return intAreaSlice, nil
}

// Part1 solves the first part of day 1 challenge.
func Part1(input string) int {
	intAreaSlice, err := parseInput(input)
	if err != nil {
		// Handle error
		return -1
	}

	totalSquareFeet := 0

	for _, areas := range intAreaSlice {
		l, w, h := areas[0], areas[1], areas[2]
		smallestSide := min(l*w, l*h, h*w)
		totalSquareFeet += (2 * l * w) + (2 * l * h) + (2 * w * h) + smallestSide
	}

	return totalSquareFeet
}

func Part2(input string) int {

	intAreaSlice, err := parseInput(input)
	if err != nil {
		// Handle error
		return -1
	}

	totalRibbonBowLength := 0

	for _, areas := range intAreaSlice {
		l, w, h := areas[0], areas[1], areas[2]
		smallestPerimeter := min(l*2+w*2, l*2+h*2, h*2+w*2)
		totalRibbonBowLength += l*w*h + smallestPerimeter
	}

	return totalRibbonBowLength
}
