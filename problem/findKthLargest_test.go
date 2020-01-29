package problem

import (
	"fmt"
	"sort"
	"testing"
)

func findKthLargest(nums []int, k int) int {
	stack := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			stack = append(stack, nums[0])
		} else {
			stack = append(stack, nums[i])
			if len(stack) >= k {
				sort.Ints(stack)
				stack = stack[len(stack)-k:]
			}
		}
		fmt.Println(stack)
	}
	return stack[0]
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
