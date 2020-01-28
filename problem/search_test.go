package problem

import (
	"fmt"
	"math"
	"testing"
)

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	k := int(math.Ceil(float64(l / 2)))
	if nums[k] == target {
		return k
	} else {
		left := nums[:k]
		right := nums[k:]
		leftResult := search(left, target)
		rightResult := search(right, target)
		if leftResult != -1 {
			return leftResult
		}
		if rightResult != -1 {
			return k + rightResult
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
