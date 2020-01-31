package problem

import (
	"fmt"
	"testing"
)

func nextPermutation(nums []int) []int {
	i := len(nums) - 1
	for i > 0 {
		fmt.Println(nums)
		if nums[i] > nums[i-1] {
			//查找一个比i-1 还大的数
			for j := len(nums) - 1; j >= i; j-- {
				if nums[i-1] < nums[j] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					break
				}
			}
			break
		} else {
			i--
		}
	}
	revertNums(nums, i)
	return nums
}
func revertNums(nums []int, start int) []int {
	for i, j := start, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
func TestNextPermutation(t *testing.T) {
	fmt.Println(nextPermutation([]int{4, 2, 0, 2, 3, 2, 0}))
}
