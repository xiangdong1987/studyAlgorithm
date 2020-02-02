package problem

func spiralOrder(matrix [][]int) []int {
	status := 1
	v := len(matrix)
	if v == 0 {
		return []int{}
	}
	h := len(matrix[0])
	if v == 1 {
		return matrix[0]
	}
	mapPath := make(map[int][]int, v)
	for i := 0; i < v; i++ {
		mapPath[i] = make([]int, h)
	}
	i := 0
	j := 0
	var result []int
	for i < v {
		for j < h {
			if len(result) == v*h {
				return result
			}
			mapPath[i][j] = 1
			if status == 1 {
				if j != h-1 && mapPath[i][j+1] != 1 {
					result = append(result, matrix[i][j])
					j++
				} else {
					status = 2
					result = append(result, matrix[i][j])
					i++
					continue
				}
			}
			if status == 2 {
				if i != v-1 && mapPath[i+1][j] != 1 {
					result = append(result, matrix[i][j])
					i++
				} else {
					status = 3
					result = append(result, matrix[i][j])
					j--
					continue
				}
			}
			if status == 3 {
				if j != 0 && mapPath[i][j-1] != 1 {
					result = append(result, matrix[i][j])
					j--
				} else {
					status = 4
					result = append(result, matrix[i][j])
					i--
					continue
				}

			}
			if status == 4 {
				if i != 0 && mapPath[i-1][j] != 1 {
					result = append(result, matrix[i][j])
					i--
				} else {
					status = 1
					result = append(result, matrix[i][j])
					j++
					continue
				}
			}
		}
	}
	return result
}
