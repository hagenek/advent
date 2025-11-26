package day9

import (
	"slices"
	"strconv"
	"strings"

	"github.com/hagenek/advent/2015/golang/utils"
)

func makeKey(word1, word2 string) string {
	if word1 > word2 {
		word1, word2 = word2, word1 // swap
	}
	return word1 + "," + word2
}

func minAndMax(input string) (int, int) {
	distances := make(map[string]int)
	cities := make(map[string]bool)

	// How to best model this?
	// Build route from scrambling the 3 places in unique orderings
	// then check distance by looking up each distance in a map?
	// one map for each city seems overkill? A list of maps, or alphabethized
	// connection map DublinLondon as key i and i+1, while i < len(route) - 1

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		words := strings.Fields(line)
		city1 := words[0]
		city2 := words[2]
		dist, err := strconv.Atoi(words[4])
		if err != nil {
			panic("Unable to parse int in string")
		}
		distanceKey := makeKey(city1, city2)
		distances[distanceKey] = dist
		cities[city1] = true
		cities[city2] = true
	}

	citiesSlice := make([]string, 0, len(cities))

	for key := range cities {
		citiesSlice = append(citiesSlice, key)
	}

	// Make or explore all possibe combinations

	permutations := utils.HeapsPermutation(citiesSlice)

	routeDistances := make([]int, len(permutations))

	for routeIndex, p := range permutations {
		routeDistance := 0
		for i, city := range p {
			if i == len(p)-1 {
				continue
			}
			cityKey := makeKey(city, p[i+1])
			routeDistance += distances[cityKey]
		}
		routeDistances[routeIndex] = routeDistance
	}

	return slices.Min(routeDistances), slices.Max(routeDistances)

}

// Part1 solves Advent of Code 2015 Day 9, Part 1.
func Part1(input string) int {
	min, _ := minAndMax(input)
	return min
}

// Part2 solves Advent of Code 2015 Day 9, Part 2.
func Part2(input string) int {
	// TODO: implement solution
	_, max := minAndMax(input)
	return max
}
