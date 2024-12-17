package solutions

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
