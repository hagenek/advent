package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day1Part1(input string) int {
	x, y := 0, 0
	moves := strings.Split(input, ", ")
	currentDir := 0

	for _, move := range moves {
		dir := move[0:1]
		dist, err := strconv.Atoi(move[1:])

		if err != nil {
			fmt.Printf("Unable to parse int from string %s", move[1:])
		}

		if dir == "R" {
			currentDir = (currentDir + 1) % 4
		} else {
			currentDir = (currentDir + 3) % 4
		}

		switch currentDir {
		case 0:
			x = x + dist
		case 1:
			y = y + dist
		case 2:
			x = x - dist
		case 3:
			y = y - dist
		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func Day1Part2(input string) int {
	x, y := 0, 0
	moves := strings.Split(input, ", ")
	currentDir := 0
	visited := make(map[[2]int]bool)

	for _, move := range moves {
		dir := move[0:1]
		dist, err := strconv.Atoi(move[1:])

		if err != nil {
			fmt.Printf("Unable to parse int from string %s", move[1:])
		}

		if dir == "R" {
			currentDir = (currentDir + 1) % 4
		} else {
			currentDir = (currentDir + 3) % 4
		}

		for i := 0; i < dist; i++ {
			switch currentDir {
			case 0:
				x++
			case 1:
				y++
			case 2:
				x--
			case 3:
				y--
			}
			pos := [2]int{x, y}
			if visited[pos] {
				return int(math.Abs(float64(x)) + math.Abs(float64(y)))
			}
			visited[pos] = true
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func Day1() (int, int) {
	input := readInputFile(1)

	return Day1Part1(input), Day1Part2(input)
}
