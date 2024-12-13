package days

import (
	"fmt"
	"strings"
)

var Day2Solutions = struct {
	Part1 func(string) int
	Part2 func(string) int
}{
	Part1: part1,
	Part2: part2,
}

func parseToProgram(input string) []int {
	var program []int
	parts := strings.Split(input, ",")
	for _, val := range parts {
		var num int
		fmt.Sscanf(val, "%d", &num)
		program = append(program, num)
	}
	return program
}

type Instructions struct {
	opcode int
	x      int
	y      int
	res_i  int
}

func do_operation(program *[]int, inst Instructions) {
	switch inst.opcode {
	case 1:
		(*program)[inst.res_i] = (*program)[inst.x] + (*program)[inst.y]
	case 2:
		(*program)[inst.res_i] = (*program)[inst.x] * (*program)[inst.y]
	default:
		break
	}
}

func part1(input string) int {
	fmt.Printf("Receiving input %s", input)
	p := parseToProgram(input)

	p[1] = 12
	p[2] = 2

	fmt.Printf("After 1202 init: %v\n", p[:10]) // print first 10 numbers to verify

	fmt.Printf("Parsed program %v", p)

	for i := 0; i < len(p); i += 4 {
		if p[i] == 99 {
			break
		}

		if i+3 > len(p)-1 {
			break
		}

		do_operation(&p, Instructions{
			opcode: p[i],
			x:      p[i+1],
			y:      p[i+2],
			res_i:  p[i+3],
		})
	}

	return p[0]
}

func part2(input string) int {
	program := parseToProgram(input)
	return program[0]
}

func Day02(input string) (int, int) {
	return part1(input), part2(input)
}
