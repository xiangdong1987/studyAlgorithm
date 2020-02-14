package problem

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	pre := nums[0]
	n := 1
	for len(nums) != n {
		if nums[n] != pre {
			pre = nums[n]
			n++
		} else {
			nums = append(nums[:n-1], nums[n:]...)
		}
	}
	return n
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
