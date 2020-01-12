package problem

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var result [][]int
	if len(nums) < 3 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println(i)
		k := i + 1
		j := len(nums) - 1
		for k < j {
			if nums[i]+nums[j]+nums[k] == 0 {
				fmt.Println(i, k, j)
				result = append(result, []int{nums[i], nums[k], nums[j]})
				//去除左侧重复
				for k < j && nums[k] == nums[k+1] {
					k++
				}
				//去除用测重复
				for k < j && nums[j] == nums[j-1] {
					j--
				}
				//下一组
				k++
				j--
			} else if nums[i]+nums[j]+nums[k] > 0 {
				j--
			} else {
				k++
			}
		}
	}
	return result
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
