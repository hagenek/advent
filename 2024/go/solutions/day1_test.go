package solutions

import "testing"

func TestDay1Part1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	want := 11
	got := Day1Part1(input)
	if got != want {
		t.Errorf("Day1Part1() = %v, want %v", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	want := 31
	got := Day1Part2(input)
	if got != want {
		t.Errorf("Day1 Part2() = %v, want %v", got, want)
	}
}
