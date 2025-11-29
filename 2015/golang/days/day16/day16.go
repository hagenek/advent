package day16

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	sueRe      = regexp.MustCompile(`Sue (\d+): (.+)`)
	propRe     = regexp.MustCompile(`(\w+): (\d+)`)
	masterTape = map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
)

type matchFunc func(key string, expected, actual int) bool

func exactMatch(_ string, expected, actual int) bool {
	return expected == actual
}

func rangeMatch(key string, expected, actual int) bool {
	switch key {
	case "cats", "trees":
		return actual > expected
	case "pomeranians", "goldfish":
		return actual < expected
	default:
		return expected == actual
	}
}

func findSue(mainMap map[string]int, input string, matcher matchFunc) int {
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		m := sueRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		n, _ := strconv.Atoi(m[1])
		props := propRe.FindAllStringSubmatch(m[2], -1)

		isMatch := true
		for _, kv := range props {
			key := kv[1]
			value, _ := strconv.Atoi(kv[2])

			expected, exists := mainMap[key]
			if !exists {
				continue
			}

			if !matcher(key, expected, value) {
				isMatch = false
				break
			}
		}

		if isMatch {
			return n
		}
	}
	return -1
}

func Part1(input string) int {
	return findSue(masterTape, input, exactMatch)
}

func Part2(input string) int {
	return findSue(masterTape, input, rangeMatch)
}
