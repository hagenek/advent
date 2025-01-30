package solutions

import (
	"fmt"
	"strings"
)

// Guard goes 1 and 1 ++ in direction untill next space has a #
// then turns 90 degrees right (changes direction) and starts going again
// while guard walks, add positions visited to a set (go map)? and then return len (size) of that map
// to solve the task

func addPosAndDir(pos [2]int, dir [2]int) [2]int {
	return [2]int{pos[0] + dir[0], pos[1] + dir[1]}
}

func moveGuard(grid [][]rune, currPos [2]int, dir [2]int) ([2]int, [2]int) {
	nextPos := addPosAndDir(currPos, dir)

	if nextPos[0] < 0 || nextPos[0] >= len(grid) ||
		nextPos[1] < 0 || nextPos[1] >= len(grid[0]) {
		return currPos, dir
	}

	if grid[nextPos[0]][nextPos[1]] == '#' {
		return currPos, turnRight(dir)
	}

	return nextPos, dir
}

func turnRight(dir [2]int) [2]int {
	return [2]int{dir[1], -dir[0]}
}

func findPositionOfGuard(grid [][]rune) ([2]int, [2]int) {
	for ix := range grid {
		for iy := range grid[ix] {
			switch grid[ix][iy] {
			case '^':
				return [2]int{ix, iy}, [2]int{-1, 0}
			case 'v':
				return [2]int{ix, iy}, [2]int{1, 0}
			case '<':
				return [2]int{ix, iy}, [2]int{0, -1}
			case '>':
				return [2]int{ix, iy}, [2]int{0, 1}
			}
		}
	}
	panic("Found no guard!!")
}

func Day6Part1(input string) int {
	lines := strings.Split(input, "\n")
	grid := LinesToRuneGrid(lines)
	guardPos, guardDir := findPositionOfGuard(grid)

	pos := guardPos
	moveHistory := make(map[[2]int]bool)
	dir := guardDir
	moveHistory[guardPos] = true

	moveCount := 0
	for {
		newPos, newDir := moveGuard(grid, pos, dir)
		moveCount++

		if pos == newPos && dir == newDir {
			break
		}

		moveHistory[newPos] = true
		pos = newPos
		dir = newDir
	}

	return len(moveHistory)
}

func fourthPointPossible(x Point, y Point, z Point, grid Grid) bool {
	dx := y.Sub(x)
	w := z.Add(dx)

	return grid.InBounds(w) && grid[w.Row][w.Col] != '#'
}

func Day6Part2(input string) int {
	lines := strings.Split(input, "\n")
	grid := NewGridFromStrings(lines)
	guardPos, _ := findPositionOfGuard(grid)

	pos := guardPos
	nwPos := Point{pos[0] - 1, pos[1]}

	var nePos Point
	neFound := false

	var swPos Point
	swFound := false

	var sePos Point

	// Find first wall to the right of guard
	for index, rune := range grid[pos[0]][pos[1]:] {
		if rune == '#' {
			nePos = Point{pos[0], pos[1] + index}
			neFound = true
			break
		}
	}

	if !neFound {
		fmt.Printf("Debug: neFound is false\n")
		row := nePos.Row
		col := nePos.Col - 1
		fmt.Printf("Debug: row=%d, col=%d\n", row, col)
		if col >= 0 {
			fmt.Printf("Debug: searching for point from {%d,%d} going South\n", row, nePos.Col-1)
			point, ok := grid.FindFirst(Point{row, nePos.Col - 1}, South, '#')
			if ok {
				fmt.Printf("Debug: found point at {%d,%d}\n", point.Row, point.Col)
				swFound = true
				swPos = point
			} else {
				fmt.Printf("Debug: no point found, returning 0\n")
				return 0
			}
		}
	}

	if swFound && neFound {
		if fourthPointPossible(swPos, nePos, nwPos, grid) {
			return 1
		} else {
			return 0
		}
	} else {
		// try to find south east
		if neFound {
			point, ok := grid.FindFirst(nePos, South, '#')
			if ok {
				if point.Col == swPos.Col && point.Row > swPos.Row {
					sePos = point
					if fourthPointPossible(sePos, nePos, nwPos, grid) {
						return 1
					} else {
						return 0
					}
				}
			}
		} else if swFound {
			point, ok := grid.FindFirst(swPos, East, '#')
			if ok {
				if point.Col == nePos.Col && point.Row > nePos.Row {
					sePos = point
					if fourthPointPossible(sePos, swPos, nwPos, grid) {
						return 1
					} else {
						return 0
					}
				}
			}
		} else {
			return 0
		}
	}
	return -1
}
func Day6() (int, int) {
	inp := readInputFile(6)
	return Day6Part1(inp), Day6Part2(inp)
}
