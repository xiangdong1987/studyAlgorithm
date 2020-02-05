package problem

import (
	"fmt"
	"sort"
	"testing"
)

func sortedMerge(nums1 []int, m int, nums2 []int, n int) []int {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] >= nums2[j] {
				fmt.Println(i, j)
				nums1[i], nums2[j] = nums2[j], nums1[i]
				sort.Ints(nums2)
				break
			} else {
				continue
			}
		}
	}
	nums1 = append(nums1[:m], nums2...)
	return nums1
}

func TestSortedMerge(t *testing.T) {
	fmt.Println(sortedMerge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
