package problem

func xorQueries(arr []int, queries [][]int) []int {
	var result []int
	for _, value := range queries {
		result = append(result, getValue(value[0], value[1], arr))
	}
	return result
}

func dictionary(arr []int) [][]int {
	var result [][]int
	for i := 0; i < len(arr); i++ {
		result = append(result, make([]int, len(arr)))
		for j := i; j < len(arr); j++ {
			if j == i {
				result[i][j] = arr[i]
			} else {
				result[i][j] = result[i][j-1] ^ arr[j]
			}
		}
	}
	return result
}

func getValue(start, end int, arr []int) int {
	var value int
	value = arr[start]
	for i := start + 1; i <= end; i++ {
		value = value ^ arr[i]
	}
	return value
}
