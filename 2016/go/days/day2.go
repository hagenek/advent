package days

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	up    *Node
	down  *Node
	left  *Node
	right *Node
	value string
}

func (n *Node) Up() *Node {
	if n.up == nil {
		return n
	}
	return n.up
}

func (n *Node) Down() *Node {
	if n.down == nil {
		return n
	}
	return n.down
}

func (n *Node) Left() *Node {
	if n.left == nil {
		return n
	}
	return n.left
}

func (n *Node) Right() *Node {
	if n.right == nil {
		return n
	}
	return n.right
}

func NewNumPad() *Node {
	one := &Node{value: "1"}
	two := &Node{value: "2"}
	three := &Node{value: "3"}
	four := &Node{value: "4"}
	five := &Node{value: "5"}
	six := &Node{value: "6"}
	seven := &Node{value: "7"}
	eight := &Node{value: "8"}
	nine := &Node{value: "9"}

	one.down = four
	one.right = two

	two.left = one
	two.right = three
	two.down = five

	three.left = two
	three.down = six

	four.up = one
	four.right = five
	four.down = seven

	five.up = two
	five.down = eight
	five.left = four
	five.right = six

	six.up = three
	six.left = five
	six.down = nine

	seven.up = four
	seven.right = eight

	eight.up = five
	eight.left = seven
	eight.right = nine

	nine.up = six
	nine.left = eight

	return five
}

func NewFancyNumPad() *Node {
	one := &Node{value: "1"}
	two := &Node{value: "2"}
	three := &Node{value: "3"}
	four := &Node{value: "4"}
	five := &Node{value: "5"}
	six := &Node{value: "6"}
	seven := &Node{value: "7"}
	eight := &Node{value: "8"}
	nine := &Node{value: "9"}
	a := &Node{value: "A"}
	b := &Node{value: "B"}
	c := &Node{value: "C"}
	d := &Node{value: "D"}

	one.down = three

	two.right = three
	two.down = six

	three.left = two
	three.right = four
	three.up = one
	three.down = seven

	four.left = three
	four.down = eight

	five.right = six

	six.left = five
	six.right = seven
	six.up = two
	six.down = a

	seven.left = six
	seven.right = eight
	seven.up = three
	seven.down = b

	eight.left = seven
	eight.right = nine
	eight.up = four
	eight.down = c

	nine.left = eight

	a.right = b
	a.up = six

	b.left = a
	b.right = c
	b.up = seven
	b.down = d

	c.left = b
	c.up = eight

	d.up = b

	return five
}

func Day2Part1(input string) int {
	lines := strings.Split(input, "\n")
	numpad := NewNumPad()
	code := ""

	for _, line := range lines {
		fmt.Println(line)
		for _, move := range line {
			switch move {
			case 'L':
				numpad = numpad.Left()
			case 'R':
				numpad = numpad.Right()
			case 'U':
				numpad = numpad.Up()
			case 'D':
				numpad = numpad.Down()
			}
		}
		code = code + numpad.value
	}
	intCode, _ := strconv.Atoi(code)
	return intCode
}

func Day2Part2(input string) string {
	lines := strings.Split(input, "\n")
	numpad := NewFancyNumPad()
	code := ""

	for _, line := range lines {
		fmt.Println(line)
		for _, move := range line {
			switch move {
			case 'L':
				numpad = numpad.Left()
			case 'R':
				numpad = numpad.Right()
			case 'U':
				numpad = numpad.Up()
			case 'D':
				numpad = numpad.Down()
			}
		}
		fmt.Printf("Numpad node: %+v\n", numpad)
		code = code + numpad.value
	}
	return code
}

func Day2() (int, string) {
	input := readInputFile(2)

	return Day2Part1(input), Day2Part2(input)
}
