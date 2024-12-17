package daysIdiomatic

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Direction represents cardinal directions
type Direction int

const (
	North Direction = iota
	South
	East
	West
)

// Position represents a 2D coordinate
type Position struct {
	x, y int
}

// Manhattan returns the Manhattan distance from origin (0,0)
func (p Position) Manhattan() int {
	return abs(p.x) + abs(p.y)
}

// Move updates the position based on current facing direction and distance
func (p *Position) Move(facing Direction, distance int) {
	switch facing {
	case North:
		p.y += distance
	case South:
		p.y -= distance
	case East:
		p.x += distance
	case West:
		p.x -= distance
	}
}

type Navigation struct {
	pos     Position
	facing  Direction
	visited map[Position]bool
}

func NewNavigation() *Navigation {
	return &Navigation{
		pos:     Position{0, 0},
		facing:  North,
		visited: make(map[Position]bool),
	}
}

func (n *Navigation) Turn(instruction byte) error {
	switch instruction {
	case 'L':
		n.facing = Direction((int(n.facing) + 3) % 4)
	case 'R':
		n.facing = Direction((int(n.facing) + 1) % 4)
	default:
		return fmt.Errorf("invalid turn instruction: %c", instruction)
	}
	return nil
}

func (n *Navigation) ProcessInstruction(instruction string) (*Position, error) {
	instruction = strings.TrimSpace(instruction)
	if len(instruction) < 2 {
		return nil, fmt.Errorf("invalid instruction format: %s", instruction)
	}

	if err := n.Turn(instruction[0]); err != nil {
		return nil, err
	}

	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid distance value: %s", instruction[1:])
	}

	// For Part 2: Track each position visited
	var firstDuplicate *Position
	for i := 0; i < distance; i++ {
		n.pos.Move(n.facing, 1)

		// Check for first duplicate position (Part 2)
		if n.visited[n.pos] && firstDuplicate == nil {
			dup := n.pos // Create a copy
			firstDuplicate = &dup
		}
		n.visited[n.pos] = true
	}

	return firstDuplicate, nil
}

// Day1Part1 solves Part 1 of the puzzle
func Day1Part1(input string) (int, error) {
	nav := NewNavigation()
	instructions := strings.Split(input, ",")

	for _, instruction := range instructions {
		if _, err := nav.ProcessInstruction(instruction); err != nil {
			return 0, fmt.Errorf("processing instruction %s: %w", instruction, err)
		}
	}

	return nav.pos.Manhattan(), nil
}

// Day1Part2 solves Part 2 of the puzzle
func Day1Part2(input string) (int, error) {
	return 0, nil
}

// Day1 reads the input file and returns solutions for both parts
func Day1() (int, int, error) {
	content, err := os.ReadFile("inputs/day01.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("reading input file: %w", err)
	}

	inputString := string(content)
	part1, err := Day1Part1(inputString)
	if err != nil {
		log.Printf("Error in Part 1: %v", err)
		part1 = 0
	}

	part2, err := Day1Part2(inputString)
	if err != nil {
		log.Printf("Error in Part 2: %v", err)
		part2 = 0
	}

	return part1, part2, nil
}

// abs returns the absolute value of x
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
