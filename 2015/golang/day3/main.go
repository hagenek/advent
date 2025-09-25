package main

import (
	"fmt"
	"os"
	"log"
)

// Position represents a house location on the 2D grid
type Position struct {
	x int
	y int
}

func main() {
	// Read input from file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	uniqueHouses := countUniqueHouses(string(input))
	fmt.Printf("Number of houses receiving at least one present: %d\n", uniqueHouses)
}

func countUniqueHouses(directions string) int {
	// Map to track visited houses
	visited := make(map[Position]bool)
	
	// Santa starts at origin and delivers a present
	currentPos := Position{0, 0}
	visited[currentPos] = true
	
	// Process each direction
	for _, move := range directions {
		switch move {
		case '^': // North
			currentPos.y++
		case 'v': // South
			currentPos.y--
		case '>': // East
			currentPos.x++
		case '<': // West
			currentPos.x--
		}
		
		// Mark current position as visited
		visited[currentPos] = true
	}
	
	// Return count of unique houses visited
	return len(visited)
}