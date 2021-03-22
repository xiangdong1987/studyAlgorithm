package problem

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	current := 0
	//选数组
	nA := len(firstList)
	nB := len(secondList)
	//求交集
	var result [][]int
	i := 0
	j := 0
	var tmp []int
	for i < nA && j < nB {
		current, tmp = interval(current, firstList[i], secondList[j])
		if tmp != nil {
			result = append(result, tmp)
		}
		if i == nA || j == nB {
			break
		}
		if firstList[i][1] > secondList[j][1] {
			j++
		} else if firstList[i][1] < secondList[j][1] {
			i++
		} else if firstList[i][1] == secondList[j][1] {
			i++
			j++
		}
	}
	return result
}

func interval(current int, a, b []int) (int, []int) {
	var start, end int
	start = current

	if start <= a[0] {
		start = a[0]
	}
	if start <= b[0] {
		start = b[0]
	}
	current = a[1]
	end = a[1]

	if current >= b[1] {
		current = b[1]
		end = b[1]
	}

	if start <= end {
		return current, []int{start, end}
	} else {
		return current, nil
	}
}
