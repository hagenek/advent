package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1Part1(input string) int {
	lines := strings.Split(input, "\n")
	freq := 0

	for _, line := range lines {
		op := line[0:1]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			return 0
		}

		if op == "-" {
			freq -= n
		} else {
			freq += n
		}
	}
	return freq
}

// Basic pattern for do-while loops in Go
// Use a for loop with a break in the middle:
//   for {
//     // do stuff
//     if !condition { break }
//   }

func Day1Part2(input string) int {
	lines := strings.Split(input, "\n")
	freq := 0
	m := make(map[int]bool)

	m[0] = true

	for {
		for _, line := range lines {
			n, err := strconv.Atoi(line[0:])
			if err != nil {
				n, err = strconv.Atoi(line[0 : len(line)-1])
				if err != nil {
					fmt.Printf("Error converting string to integer: %v\n", err)
					return 0
				}
			}

			freq += n

			if ok := m[freq]; ok {
				return freq
			} else {
				m[freq] = true
			}
		}
	}
}

func Day1() (int, int) {
	inp := readInputFile(1)

	return Day1Part1(inp), Day1Part2(inp)
}
