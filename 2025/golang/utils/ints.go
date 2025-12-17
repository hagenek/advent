package utils

import "strconv"

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func Int[T ~string | ~rune](i T) int {
	var s string
	switch v := any(i).(type) {
	case string:
		s = v
	case rune:
		s = string(v)
	default:
		return 0
	}

	if s == "" {
		return 0
	}

	n, _ := strconv.Atoi(s)
	return n
}
