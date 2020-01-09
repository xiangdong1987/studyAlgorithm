package problem

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for i, val := range nums {
		pos, can := findNum(nums, target-val, i)
		if can {
			result = append(result, i)
			result = append(result, pos)
			break
		} else {
			continue
		}
	}
	return result
}
func findNum(nums []int, target int, position int) (pos int, can bool) {
	can = false
	for i, val := range nums {
		if val == target && i != position {
			can = true
			return i, can
		}
	}
	return 0, can
}

func TestTwoSum(t *testing.T) {
	fmt.Print(twoSum([]int{3, 2, 4}, 6))
}
