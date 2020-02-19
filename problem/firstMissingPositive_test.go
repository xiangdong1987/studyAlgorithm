package problem

import (
	"fmt"
	"testing"
)

func firstMissingPositive(nums []int) int {
	//第一遍清除无效
	n := len(nums)
	if n < 1 {
		return 1
	}
	min := nums[0]
	max := nums[0]
	containOne := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			containOne = true
		}
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] > n {
			nums[i] = 1
		}
		if nums[i] <= 0 {
			nums[i] = 1
		}
	}
	fmt.Println(nums)
	if min > 1 {
		return 1
	}
	if max < 1 {
		return 1
	}
	fmt.Println(nums)
	//以索引为hash值的正负为是否存在过
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		if current <= 0 {
			current = 0 - current
		}
		if nums[current-1] > 0 {
			nums[current-1] = 0 - nums[current-1]
		}
	}
	if containOne {
		nums[0] = -1
	} else {
		nums[0] = 1
	}
	fmt.Println(nums, containOne)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func TestFirstMissingPositive(t *testing.T) {
	fmt.Println(firstMissingPositive([]int{1, 2, 2, 1, 3, 1, 0, 4, 0}))
}
