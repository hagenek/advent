package solutions

import "testing"

func TestDay2(t *testing.T) {
	t.Run("Part 1 test", func(t *testing.T) {
		input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
		result := Day2Part1(input)
		if result != 2 {
			t.Errorf("Expected 2, got %d", result)
		}
	})

	t.Run("Part 2 test", func(t *testing.T) {
		input := "test input"
		result := Day2Part2(input)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})

}
