package problem

import (
	"math"
	"testing"
)

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		max = int(math.Max(float64(nums[i]), float64(max)))
	}
	return max
}

func TestMaxSubArray(t *testing.T) {
	println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
