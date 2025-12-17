package day1

import (
	"strconv"
	"strings"
)

type Dial struct {
	pos          int
	landedOnZero int
	passedZero   int
}

func (d *Dial) RL(n int) {
	for range n {
		d.pos--
		if d.pos < 0 {
			d.pos = 99
		}
		if d.pos == 0 {
			d.passedZero++
		}
	}
}

func (d *Dial) RR(n int) {
	for range n {
		d.pos++
		if d.pos > 99 {
			d.pos = 0
		}
		if d.pos == 0 {
			d.passedZero++
		}
	}
}

func Solve(input string) (int, int) {
	dial := Dial{}
	dial.pos = 50
	for ins := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		dir := ins[0]
		n, _ := strconv.Atoi(ins[1:])

		if rune(dir) == 'R' {
			dial.RR(n)
		} else {
			dial.RL(n)
		}
		if dial.pos == 0 {
			dial.landedOnZero++
		}
	}

	return dial.landedOnZero, dial.passedZero
}

func Part1(input string) int {
	p1, _ := Solve(input)
	return p1
}

func Part2(input string) int {
	_, p2 := Solve(input)
	return p2
}
