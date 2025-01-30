package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LinesToRuneGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i := range grid {
		grid[i] = make([]rune, len(lines[i]))
	}

	for x, line := range lines {
		for y, char := range line {
			grid[x][y] = char
		}
	}
	return grid
}

func listInputFiles() {
	entries, err := os.ReadDir("inputs")
	if err != nil {
		panic(fmt.Errorf("failed to read inputs directory: %w", err))
	}

	fmt.Println("\nFiles in inputs directory:")
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Printf("  - %s\n", entry.Name())
		}
	}
	fmt.Println()
}

func readInputFile(day int) string {
	data, err := os.ReadFile(fmt.Sprintf("inputs/day%d.txt", day))
	if err != nil {
		listInputFiles()
		panic(fmt.Errorf("failed to read input for day %d: %w", day, err))
	}
	return strings.TrimSpace(string(data))
}

func parseInt(r string) int {
	num, err := strconv.Atoi(r)
	if err != nil {
		fmt.Printf("Error converting %s, err: %s", r, err)
	}
	return num
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i := range slice {
		result[i] = fn(slice[i])
	}
	return result
}

func Reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

var (
	North = Point{-1, 0}
	South = Point{1, 0}
	East  = Point{0, 1}
	West  = Point{0, -1}

	NorthEast = Point{-1, 1}
	NorthWest = Point{-1, -1}
	SouthEast = Point{1, 1}
	SouthWest = Point{1, -1}
)

type Point struct {
	Row, Col int
}

type Grid [][]rune

func NewGridFromStrings(lines []string) Grid {
	grid := make(Grid, len(lines))
	for i := range grid {
		grid[i] = []rune(lines[i])
	}
	return grid
}

func (g Grid) Size() (rows, cols int) {
	if len(g) == 0 {
		return 0, 0
	}
	return len(g), len(g[0])
}

func (g Grid) InBounds(p Point) bool {
	rows, cols := g.Size()
	return p.Row >= 0 && p.Row < rows && p.Col >= 0 && p.Col < cols
}

func (p Point) Add(v Point) Point {
	return Point{p.Row + v.Row, p.Col + v.Col}
}

func (p Point) Sub(v Point) Point {
	return Point{p.Row - v.Row, p.Col - v.Col}
}

// FindFirst searches in a direction until it finds the target rune
func (g Grid) FindFirst(start Point, direction Point, target rune) (Point, bool) {
	current := start
	for g.InBounds(current) {
		if g[current.Row][current.Col] == target {
			return current, true
		}
		current = current.Add(direction)
	}
	return Point{}, false
}
