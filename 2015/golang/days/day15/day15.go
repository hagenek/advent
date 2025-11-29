package day15

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func (i Ingredient) String() string {
	return fmt.Sprintf(`{
		name:       '%s',
		capacity:   %d,
		durability: %d,
		flavor:     %d,
		texture:    %d,
		calories:   %d
}`, i.name, i.capacity, i.durability, i.flavor, i.texture, i.calories)
}

func parseIngredient(line string) Ingredient {
	re := regexp.MustCompile(`(\w+): capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)
	matches := re.FindStringSubmatch(strings.TrimSpace(line))

	fmt.Printf("Matches: %v\n", matches)
	fmt.Printf("Line: %s\n", line)

	name := matches[1]
	capacity, _ := strconv.Atoi(matches[2])
	durability, _ := strconv.Atoi(matches[3])
	flavor, _ := strconv.Atoi(matches[4])
	texture, _ := strconv.Atoi(matches[5])
	calories, _ := strconv.Atoi(matches[6])

	return Ingredient{
		name:       name,
		capacity:   capacity,
		durability: durability,
		flavor:     flavor,
		texture:    texture,
		calories:   calories,
	}
}

func (i Ingredient) cloneAndMerge(n int) Ingredient {
	ni := Ingredient{}
	for range n {
		ni.capacity += i.capacity
		ni.durability += i.durability
		ni.flavor += i.flavor
		ni.texture += i.texture
		ni.calories += i.calories
	}
	return ni
}

func multiplyIngredients(i1, i2, i3, i4 Ingredient) Ingredient {
	return Ingredient{
		name:       i1.name + " " + i2.name + " " + i3.name + " " + i4.name,
		capacity:   i1.capacity + i2.capacity + i3.capacity + i4.capacity,
		durability: i1.durability + i2.durability + i3.durability + i4.durability,
		flavor:     i1.flavor + i2.flavor + i3.flavor + i4.flavor,
		texture:    i1.texture + i2.texture + i3.texture + i4.texture,
		calories:   i1.calories + i2.calories + i3.calories + i4.calories,
	}
}

func zin(n int) int {
	if n < 0 {
		return 0
	} else {
		return n
	}
}

func sumIngredient(i Ingredient) int {
	return zin(i.capacity) * zin(i.durability) * zin(i.flavor) * zin(i.texture)
}

func Solve(input string, part2 bool) int {
	dataLines := strings.SplitSeq(strings.TrimSpace(input), "\n")

	ingredients := []Ingredient{}

	for line := range dataLines {
		ingredient := parseIngredient(line)
		fmt.Printf("Ingredient parsed: %v\n", ingredient)
		ingredients = append(ingredients, ingredient)
	}

	i1 := ingredients[0]
	i2 := ingredients[1]
	i3 := ingredients[2]
	i4 := ingredients[3]

	sums := []int{}

	for n := range 100 {
		for n2 := range 99 {
			for n3 := range 98 {
				li1 := i1.cloneAndMerge(n)
				li2 := i2.cloneAndMerge(n2)
				li3 := i3.cloneAndMerge(n3)
				li4 := i4.cloneAndMerge(100 - (n + n2 + n3))

				i3 := multiplyIngredients(li1, li2, li3, li4)
				sum := sumIngredient(i3)
				if !part2 {
					sums = append(sums, sum)
				}
				if part2 && i3.calories == 500 {
					sums = append(sums, sum)
				}
			}
		}
	}

	return slices.Max(sums)
}

func Part1(input string) int {
	return Solve(input, false)
}

// Part2 solves Advent of Code 2015 Day 15, Part 2.
// Implement the solution and return the result as an int.
// TODO: implement
func Part2(input string) int {
	return Solve(input, true)
}
