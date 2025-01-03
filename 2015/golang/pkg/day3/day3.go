package day3

import "fmt"

type Position struct {
	X, Y int
}

func (p *Position) Move(direction rune) {
	switch direction {
	case '^':
		p.Y++
	case 'v':
		p.Y--
	case '>':
		p.X++
	case '<':
		p.X--
	default:
		fmt.Println("Invalid direction")
	}
}
func Part1(input string) int {
	p := Position{0, 0}
	visited := make(map[string]bool)
	visited[coordString(p.X, p.Y)] = true
	for _, char := range input {
		p.Move(char)
	}
	return len(visited)
}

func Part2(input string) int {
	p1 := Position{0, 0}
}

