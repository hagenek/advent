package day3

type Position struct {
	X, Y int
}

// Part1 returns the number of unique houses that receive at least one present
// based on the movement instructions in the input string.
func Part1(input string) int {
	pos := Position{0, 0}
	visited := make(map[Position]int)
	visited[pos] = 1

	for _, c := range input {
		switch c {
		case '^':
			pos.Y++
		case '>':
			pos.X++
		case '<':
			pos.X--
		case 'v':
			pos.Y--
		default:
			continue
		}
		visited[pos]++
	}

	return len(visited)
}

func Part2(input string) int {
	posSanta := Position{0, 0}
	posRobot := Position{0, 0}
	visited := make(map[Position]int)
	visited[Position{0, 0}] = 2

	for idx, c := range input {
		if idx%2 == 0 {
			switch c {
			case '^':
				posSanta.Y++
			case '>':
				posSanta.X++
			case '<':
				posSanta.X--
			case 'v':
				posSanta.Y--
			default:
				continue
			}
			visited[posSanta]++
		} else {
			switch c {
			case '^':
				posRobot.Y++
			case '>':
				posRobot.X++
			case '<':
				posRobot.X--
			case 'v':
				posRobot.Y--
			default:
				continue
			}
			visited[posRobot]++
		}
	}
	return len(visited)
}
