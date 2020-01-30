package problem

import (
	"fmt"
	"sort"
	"testing"
)

func nextPermutation(nums []int) []int {
	//判断是否逆序
	isReversal := true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			isReversal = false
		}
	}
	fmt.Println(isReversal)
	if isReversal {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		return nums
	}

	x := len(nums) - 2
	last := len(nums) - 1
	for x >= 0 {
		fmt.Println(last, nums)
		if nums[last] > nums[x] {
			nums[x], nums[last] = nums[last], nums[x]
			fmt.Println(x)
			back := nums[x+1:]
			sort.Ints(back)
			for i, j := 0, len(back)-1; i < j; i, j = i+1, j-1 {
				back[i], back[j] = back[j], back[i]
			}
			break
		} else {
			if x == 0 {
				last--
				x = last - 1
			} else {
				x--
			}
		}
	}
	return nums
}

func TestNextPermutation(t *testing.T) {
	fmt.Println(nextPermutation([]int{4, 2, 0, 2, 3, 2, 0}))
}
