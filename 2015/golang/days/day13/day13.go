package day13

import (
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/hagenek/advent/2015/golang/utils"
)

func ParseHappiness(input string) map[string]map[string]int {
	result := make(map[string]map[string]int)
	re := regexp.MustCompile(`(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+)\.`)

	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		person := matches[1]
		gainLose := matches[2]
		amount, _ := strconv.Atoi(matches[3])
		neighbor := matches[4]

		if gainLose == "lose" {
			amount = -amount
		}

		if result[person] == nil {
			result[person] = make(map[string]int)
		}
		result[person][neighbor] = amount
	}

	return result
}

func calculateMaxHappiness(happiness map[string]map[string]int) int {
	people := []string{}
	for key := range happiness {
		people = append(people, key)
	}

	permutations := utils.HeapsPermutation(people)

	scores := []int{}

	for _, p := range permutations {
		score := 0
		for i, person := range p {
			rightSide := ""
			leftSide := ""
			if i == len(p)-1 {
				leftSide = p[0]
				rightSide = p[i-1]
				score += happiness[person][leftSide] + happiness[person][rightSide]
			} else if i == 0 {
				leftSide = p[i+1]
				rightSide = p[len(p)-1]
				score += happiness[person][leftSide] + happiness[person][rightSide]
			} else {
				leftSide = p[i+1]
				rightSide = p[i-1]
				score += happiness[person][leftSide] + happiness[person][rightSide]
			}
		}
		scores = append(scores, score)
	}

	return slices.Max(scores)
}

// Part1 solves the first part of day 13 challenge.
func Part1(input string) int {
	happiness := ParseHappiness(input)
	return calculateMaxHappiness(happiness)
}

// Part2 solves the second part of day 13 challenge.
func Part2(input string) int {
	happiness := ParseHappiness(input)
	people := []string{}

	for key := range happiness {
		people = append(people, key)
	}

	for _, person := range people {
		happiness[person]["Georg"] = 0
	}
	happiness["Georg"] = make(map[string]int)
	for _, person := range people {
		happiness["Georg"][person] = 0
	}

	return calculateMaxHappiness(happiness)
}
