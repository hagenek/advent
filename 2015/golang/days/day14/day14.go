package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Reindeer struct {
	name     string
	rest     int
	restcap  int
	flying   int
	flycap   int
	speed    int
	distance int
	points   int
}

func tick(r Reindeer) Reindeer {
	if r.flying == 0 {
		if r.rest == 1 {
			r.rest = r.restcap
			r.flying = r.flycap
			return r
		}
		// regular rest tick
		if r.rest > 0 {
			r.rest--
			return r
		}
	}
	// reindeer is flying tick
	if r.flying > 0 {
		r.flying = r.flying - 1
		r.distance += r.speed
		return r
	}
	fmt.Printf("Unknown state %v+", r)
	return r
}

func initReindeer(line string) Reindeer {
	re := regexp.MustCompile(`(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds.`)
	matches := re.FindStringSubmatch(strings.TrimSpace(line))

	if matches == nil {
		panic("OH NO")
	}

	name := matches[1]
	speed, _ := strconv.Atoi(matches[2])
	flyingcap, _ := strconv.Atoi(matches[3])
	restcap, _ := strconv.Atoi(matches[4])

	return Reindeer{
		name:    name,
		speed:   speed,
		flycap:  flyingcap,
		flying:  flyingcap,
		rest:    restcap,
		restcap: restcap,
	}
}

func highestDistance(rs []Reindeer) int {
	highest := 0
	for _, r := range rs {
		if r.distance > highest {
			highest = r.distance
		}
	}
	return highest
}

func awardPointsToLeaders(rs []Reindeer) []Reindeer {
	highest := 0
	for _, r := range rs {
		if r.distance > highest {
			highest = r.distance
		}
	}
	newRs := []Reindeer{}
	for _, r := range rs {
		if r.distance == highest {
			r.points++
		}
		newRs = append(newRs, r)
	}
	return newRs
}

func highestPoints(rs []Reindeer) int {
	highest := 0
	for _, r := range rs {
		if r.points > highest {
			highest = r.points
		}
	}
	return highest
}

// Part1 solves the first part of day 14 challenge.
func Part1(input string) int {
	reindeers := []Reindeer{}
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		reindeers = append(reindeers, initReindeer(line))
	}
	oldRs := reindeers
	for range 2503 {
		newRs := []Reindeer{}
		for _, r := range oldRs {
			r = tick(r)
			newRs = append(newRs, r)
		}
		oldRs = newRs
	}

	return highestDistance(oldRs)
}

// Part2 solves the second part of day 14 challenge.
func Part2(input string) int {
	reindeers := []Reindeer{}
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		reindeers = append(reindeers, initReindeer(line))
	}
	oldRs := reindeers
	for range 2503 {
		newRs := []Reindeer{}
		for _, r := range oldRs {
			r = tick(r)
			newRs = append(newRs, r)
		}
		newRs = awardPointsToLeaders(newRs)
		oldRs = newRs
	}

	return highestPoints(oldRs)
}
