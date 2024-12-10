// days/day01.go
package days

import (
	"math"
	"strconv"
	"strings"
)

func calculateFuel(mass int) int {
	mass_div_three := mass / 3
	f64fuel := math.Round(float64(mass_div_three))
	return int(f64fuel) - 2
}

func Day1(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	totalFuel := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		mass, err := strconv.Atoi(line)
		if err != nil {
			continue // Skip invalid numbers
		}
		totalFuel += calculateFuel(mass)
	}

	return totalFuel, 0 // Part 2 not implemented yet
}
