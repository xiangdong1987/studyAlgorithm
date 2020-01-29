package problem

import (
	"math"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	prePos := 0
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			result = append(result, intervals[0])
		} else {
			pre := result[prePos]
			if pre[1] >= intervals[i][0] {
				pre[1] = int(math.Max(float64(intervals[i][1]), float64(pre[1])))
			} else {
				result = append(result, intervals[i])
				prePos++
			}
		}
	}
	return result
}
