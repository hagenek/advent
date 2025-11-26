package day7

import (
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	target    string
}

func parseLine(line string) (operation string, target string) {
	parts := strings.Split(line, " -> ")
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func getValue(token string, wires map[string]uint16) (uint16, bool) {
	// Try to parse as number
	if val, err := strconv.Atoi(token); err == nil {
		return uint16(val), true
	}
	// Try to get from wires map
	if val, ok := wires[token]; ok {
		return val, true
	}
	return 0, false
}

func processInstruction(operation string, wires map[string]uint16) (uint16, bool) {
	tokens := strings.Fields(operation)

	switch len(tokens) {
	case 1:
		// Direct assignment: "123" or "x"
		return getValue(tokens[0], wires)

	case 2:
		// NOT operation: "NOT x"
		if tokens[0] == "NOT" {
			if val, ok := getValue(tokens[1], wires); ok {
				return ^val, true
			}
		}

	case 3:
		// Binary or Shift: "x AND y" or "x LSHIFT 2"
		left, leftOk := getValue(tokens[0], wires)
		right, rightOk := getValue(tokens[2], wires)

		if !leftOk || !rightOk {
			return 0, false
		}

		switch tokens[1] {
		case "AND":
			return left & right, true
		case "OR":
			return left | right, true
		case "LSHIFT":
			return left << right, true
		case "RSHIFT":
			return left >> right, true
		}
	}

	return 0, false
}

func simulate(input string) map[string]uint16 {
	wires := make(map[string]uint16)
	instructions := []Instruction{}

	// Parse all instructions
	for line := range strings.SplitSeq(strings.TrimSpace(input), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		operation, target := parseLine(line)
		instructions = append(instructions, Instruction{operation, target})
	}

	// Keep processing until all instructions are resolved
	for len(instructions) > 0 {
		remaining := []Instruction{}

		for _, inst := range instructions {
			if value, ok := processInstruction(inst.operation, wires); ok {
				wires[inst.target] = value
			} else {
				// Can't process yet, save for next iteration
				remaining = append(remaining, inst)
			}
		}

		// If nothing was processed this round, we're stuck (shouldn't happen with valid input)
		if len(remaining) == len(instructions) {
			break
		}

		instructions = remaining
	}

	return wires
}

// Part1 solves Advent of Code 2015 Day 7, Part 1.
func Part1(input string) int {
	wires := simulate(input)
	return int(wires["a"])
}

// Part2 solves Advent of Code 2015 Day 7, Part 2.
func Part2(input string) int {
	// TODO: implement solution
	return 0
}
