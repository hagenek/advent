package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// mode
// 0 -> off
// 1 -> on
// 2 -> toggle

type Grid struct {
	cells [][]int
	w, h  int
}

type Point struct {
	x, y int
}

func NewGrid(width, height int) *Grid {
	cells := make([][]int, height)
	for i := range cells {
		cells[i] = make([]int, width)
	}
	return &Grid{cells: cells, w: width, h: height}
}

func (g *Grid) FillRect(p1 Point, p2 Point, val int) {
	for x := p1.x; x <= p2.x; x++ {
		for y := p1.y; y <= p2.y; y++ {
			g.cells[x][y] = val
		}
	}
}

func (g *Grid) ToggleRect(p1 Point, p2 Point) {
	for x := p1.x; x <= p2.x; x++ {
		for y := p1.y; y <= p2.y; y++ {
			g.cells[x][y] = 1 - g.cells[x][y]
		}
	}
}

func (g *Grid) IncrementRect(p1 Point, p2 Point, inc int) {
	for x := p1.x; x <= p2.x; x++ {
		for y := p1.y; y <= p2.y; y++ {
			g.cells[x][y] = g.cells[x][y] + inc
		}
	}
}

func (g *Grid) DecRect(p1 Point, p2 Point) {
	for x := p1.x; x <= p2.x; x++ {
		for y := p1.y; y <= p2.y; y++ {
			if g.Get(Point{x, y}) == 0 {
				continue
			}
			g.cells[x][y]--
		}
	}
}

func (g *Grid) Get(p Point) int {
	return g.cells[p.x][p.y]
}

func (g *Grid) Set(p Point, val int) {
	g.cells[p.x][p.y] = val
}

func (g *Grid) PrintGrid() {
	for y := range g.cells {
		fmt.Printf("|")
		for x := range g.cells[y] {
			switch g.Get(Point{x, y}) {
			case 0:
				fmt.Printf("  ")
			case 1:
				fmt.Printf("* ")
			}
		}
		fmt.Printf("|\n")
	}
}

func parseLine(line string) (string, Point, Point) {
	r := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	matches := r.FindStringSubmatch(line)
	if len(matches) != 6 {
		panic(fmt.Sprintf("invalid input: %s", line))
	}

	operation := matches[1]
	x1, _ := strconv.Atoi(matches[2])
	y1, _ := strconv.Atoi(matches[3])
	x2, _ := strconv.Atoi(matches[4])
	y2, _ := strconv.Atoi(matches[5])

	return operation, Point{x1, y1}, Point{x2, y2}
}

func (g *Grid) countLights() int {
	count := 0
	for x := range g.cells {
		for y := range len(g.cells) {
			count += g.cells[x][y]
		}
	}
	return count
}

// Part1 solves Advent of Code 2015 Day 6, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement
func Part1(input string) int {
	g := NewGrid(1000, 1000)

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		operation, start, end := parseLine(line)

		switch operation {
		case "turn on":
			g.FillRect(start, end, 1)
		case "turn off":
			g.FillRect(start, end, 0)
		case "toggle":
			g.ToggleRect(start, end)
		}
	}

	return g.countLights()
}

// Part2 solves Advent of Code 2015 Day 6, Part 2.
// Implement the solution and return the result as an int.
func Part2(input string) int {
	g := NewGrid(1000, 1000)

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		operation, start, end := parseLine(line)

		switch operation {
		case "turn on":
			g.IncrementRect(start, end, 1)
		case "turn off":
			g.DecRect(start, end)
		case "toggle":
			g.IncrementRect(start, end, 2)
		}
	}

	return g.countLights()
}
