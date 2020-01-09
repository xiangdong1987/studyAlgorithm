package problem

import (
	"fmt"
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var pos int
	var result float64
	if len(nums2) > 0 {
		for len(nums1) > 0 {
			if nums1[0] >= nums2[len(nums2)-1] {
				nums2 = append(nums2, nums1...)
				break
			} else {
				//取出数字插入nums2
				tmp := nums1[0]
				nums1 = append(nums1[:0], nums1[1:]...)
				nums2 = insertSlice(tmp, nums2)
			}
		}
	} else {
		nums2 = nums1
	}
	if len(nums2)%2 == 0 {
		pos = len(nums2) / 2
		result = float64(nums2[pos-1]+nums2[pos]) / 2
	} else {
		pos := int(math.Floor(float64(len(nums2) / 2)))
		result = float64(nums2[pos])
	}
	return result
}
func insertSlice(num int, nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if num <= nums[i] {
			if i == 0 {
				nums = append([]int{num}, nums...)
			}
			continue
		} else {
			nums = append(nums[:i+1], append([]int{num}, nums[i+1:]...)...)
			break
		}
	}
	return nums
}
func TestFindMedianSortedArrays(t *testing.T) {
	fmt.Print(findMedianSortedArrays([]int{2}, []int{}))
}
