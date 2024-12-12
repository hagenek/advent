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

func calculateFuelOnFuel(mass int) int {
	new_fuel := calculateFuel(mass)
	if calculateFuel(mass) <= 0 {
		return 0
	} else {
		return new_fuel + calculateFuelOnFuel(new_fuel)
	}
}

func Day1(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	totalFuelPart1 := 0
	totalFuelPart2 := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		mass, err := strconv.Atoi(line)
		if err != nil {
			continue // Skip invalid numbers
		}
		totalFuelPart1 += calculateFuel(mass)
		totalFuelPart2 += calculateFuelOnFuel(mass)
	}

	return totalFuelPart1, totalFuelPart2 // Part 2 not implemented yet
}
