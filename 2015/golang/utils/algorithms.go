package utils

func HeapsPermutation(arr []string) [][]string {
	n := len(arr)

	if n > 10 {
		return nil
	}

	var result [][]string
	var generate func(int)

	generate = func(n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
			return
		}

		for i := range n {
			generate(n - 1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}
	generate(len(arr))
	return result
}
