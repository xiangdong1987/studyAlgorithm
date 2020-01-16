package problem

import (
	"fmt"
	"testing"
)

func jump(nums []int) int {
	l := len(nums)
	j := l - 1
	if nums[0] >= l {
		return 1
	}
	if l < 2 {
		return l - 1
	}
	dp := make(map[int]int)
	dp[j] = 0
	for j >= 0 {
		if nums[j] == 0 {
			dp[j] = dp[j+1] + 1
			j--
			continue
		}
		if nums[j] >= l-1-j {
			dp[j] = 1
			for x := j; x < l-1; x++ {
				dp[x] = 1
			}
		} else {
			min := dp[j+1]
			for k := j; k < j+nums[j]; k++ {
				if dp[k+1] < min {
					min = dp[k+1]
				}
			}
			dp[j] = min + 1
		}
		j--
	}
	fmt.Println(dp)
	return dp[0]
}

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{5, 8, 1, 8, 9, 8, 7, 1, 7, 5, 8, 6, 5, 4, 7, 3, 9, 9, 0, 6, 6, 3, 4, 8, 0, 5, 8, 9, 5, 3, 7, 2, 1, 8, 2, 3, 8, 9, 4, 7, 6, 2, 5, 2, 8, 2, 7, 9, 3, 7, 6, 9, 2, 0, 8, 2, 7, 8, 4, 4, 1, 1, 6, 4, 1, 0, 7, 2, 0, 3, 9, 8, 7, 7, 0, 6, 9, 9, 7, 3, 6, 3, 4, 8, 6, 4, 3, 3, 2, 7, 8, 5, 8, 6, 0}))
	fmt.Println(jump([]int{15, 15, 15, 14, 14, 14, 14, 15, 14, 14, 13, 13, 13, 13, 13, 13, 12, 12, 12, 12, 12, 12, 12, 11, 11, 11, 11, 11, 11, 11, 10, 12, 11, 10, 11, 10, 10, 9, 10, 9, 9, 10, 9, 9, 8, 9, 8, 8, 8, 8, 8, 7, 8, 7, 7, 7, 7, 6, 7, 7, 8, 7, 6, 6, 6, 5, 5, 6, 5, 5, 4, 4, 4, 4, 4, 4, 3, 3, 3, 4, 3, 3, 3, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1}))
}
