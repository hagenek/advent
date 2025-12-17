package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// max len = 10
// seq is 2 or more digits
// naive approach build up map of sequences, sliding window style
// if found in map, return true.
// handle even len first
func hasRepeatingSequencePart1(n int) bool {
	s := strconv.Itoa(n)

	if len(s)%2 == 0 {
		s1 := s[0 : len(s)/2]
		s2 := s[len(s)/2:]
		if s1 == s2 {
			return true
		}
	} else {
		return false
	}
	return false
}

func hasRepeatingSequencePart2(n int) bool {
	s := strconv.Itoa(n)

	nOfSame := 0
	// all repeating check
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			nOfSame++
		}
	}
	if nOfSame == len(s)-1 {
		fmt.Printf("%s has only same digits", s)
		return true
	}

	// Check all possible divisors for lengths 4, 6, 8, 10
	switch len(s) {
	case 2:
		if s[0] == s[1] {
			return true
		}
	case 3:
		if s[0] == s[1] && s[1] == s[2] {
			return true
		}
	case 4:
		// Can split into 2 pieces of length 2
		s1 := s[0:2]
		s2 := s[2:4]
		if s1 == s2 {
			return true
		}
	case 6:
		// Can split into 2 pieces of length 3
		s1 := s[0:3]
		s2 := s[3:6]
		if s1 == s2 {
			return true
		}
		// Can split into 3 pieces of length 2
		s1 = s[0:2]
		s2 = s[2:4]
		s3 := s[4:6]
		if s1 == s2 && s2 == s3 {
			return true
		}
	case 8:
		// Can split into 2 pieces of length 4
		s1 := s[0:4]
		s2 := s[4:8]
		if s1 == s2 {
			return true
		}
		// Can split into 4 pieces of length 2
		s1 = s[0:2]
		s2 = s[2:4]
		s3 := s[4:6]
		s4 := s[6:8]
		if s1 == s2 && s2 == s3 && s3 == s4 {
			return true
		}
	case 9:
		s1 := s[0:3]
		s2 := s[3:6]
		s3 := s[6:9]
		if s1 == s2 && s2 == s3 {
			return true
		}

	case 10:
		// Can split into 2 pieces of length 5
		s1 := s[0:5]
		s2 := s[5:10]
		if s1 == s2 {
			return true
		}
		// Can split into 5 pieces of length 2
		s1 = s[0:2]
		s2 = s[2:4]
		s3 := s[4:6]
		s4 := s[6:8]
		s5 := s[8:10]
		if s1 == s2 && s2 == s3 && s3 == s4 && s4 == s5 {
			return true
		}
	}

	return false
}

func hasRepeatingSequencePart3(n int) bool {
	s := strconv.Itoa(n)

	// Check if all digits are the same
	allSame := true
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return true
	}

	// Check for any repeating pattern
	for patternLen := 1; patternLen <= len(s)/2; patternLen++ {
		if len(s)%patternLen == 0 {
			pattern := s[0:patternLen]
			isRepeating := true
			for i := patternLen; i < len(s); i += patternLen {
				if s[i:i+patternLen] != pattern {
					isRepeating = false
					break
				}
			}
			if isRepeating {
				return true
			}
		}
	}

	return false
}

// Part1 solves Advent of Code 2025 Day 2, Part 1.
// Implement the solution and return the result as an int.
// TODO: implement
func Part1(input string) int {
	// TODO: implement part 1

	fmt.Printf("Len of input: %d", len(input))

	idSum := 0

	for r := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		re := regexp.MustCompile(`(\d+)-(\d+)`)
		matches := re.FindStringSubmatch(r)
		fmt.Println(matches)
		startOfRange, _ := strconv.Atoi(matches[1])
		endOfRange, _ := strconv.Atoi(matches[2])

		for i := startOfRange; i <= endOfRange; i++ {
			if hasRepeatingSequencePart1(i) {
				idSum += i
			}
		}
	}

	correctSum := 1227775554
	diff := idSum - correctSum

	fmt.Printf("Should be %d, but is %d, difference %d \n", correctSum, idSum, diff)
	return idSum
}

func Part2(input string) int {
	idSum := 0

	for r := range strings.SplitSeq(strings.TrimSpace(input), ",") {
		re := regexp.MustCompile(`(\d+)-(\d+)`)
		matches := re.FindStringSubmatch(r)
		fmt.Println(matches)
		startOfRange, _ := strconv.Atoi(matches[1])
		endOfRange, _ := strconv.Atoi(matches[2])

		for i := startOfRange; i <= endOfRange; i++ {
			if hasRepeatingSequencePart2(i) {
				idSum += i
			}
		}
	}

	correctSum := 4174379265
	diff := idSum - correctSum

	fmt.Printf("Should be %d, but is %d, difference %d \n", correctSum, idSum, diff)
	return idSum
}
