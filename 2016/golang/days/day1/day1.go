package day1

import (
	"strconv"
	"strings"

	"github.com/hagenek/advent/2016/golang/utils"
)

var dirs = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func solve(input string) (int, int) {
	x, y, dir := 0, 0, 0
	visited := map[[2]int]bool{{0, 0}: true}
	part2 := -1

	for ins := range strings.SplitSeq(strings.TrimSpace(input), ", ") {
		if ins[0] == 'L' {
			dir = (dir + 3) % 4
		} else {
			dir = (dir + 1) % 4
		}
		steps, _ := strconv.Atoi(ins[1:])

		for range steps {
			x += dirs[dir][0]
			y += dirs[dir][1]
			if part2 == -1 && visited[[2]int{x, y}] {
				part2 = utils.Abs(x) + utils.Abs(y)
			}
			visited[[2]int{x, y}] = true
		}
	}
	return utils.Abs(x) + utils.Abs(y), part2
}

func Part1(input string) int {
	p1, _ := solve(input)
	return p1
}

func Part2(input string) int {
	_, p2 := solve(input)
	return p2
}
