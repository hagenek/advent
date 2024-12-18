package leetcode

func TwoSum(nums []int, t int) []int {
	m := make(map[int]int)

	for i, n := range nums {
		if _, ok := m[t-n]; ok {
			return []int{m[t-n], i}
		}
		m[n] = i
	}
	return []int{0, 0}
}
