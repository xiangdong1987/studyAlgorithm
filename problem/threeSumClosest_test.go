package problem

import (
	"fmt"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	//先排序
	sort.Ints(nums)
	fmt.Println(nums)
	//先假定最接近距离是
	closest := nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3] - target
	result := nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
	if closest == 0 {
		return result
	}
	//再通过双指针定位查找最接近的数
	for i := 0; i < len(nums); i++ {
		L := i + 1
		R := len(nums) - 1
		for L < R {
			//fmt.Println(nums[i]+nums[L]+nums[R], closest, result)
			if nums[i]+nums[L]+nums[R] == target {
				return nums[i] + nums[L] + nums[R]
			}
			if nums[i]+nums[L]+nums[R] > target {
				tmp := nums[i] + nums[L] + nums[R] - target
				if tmp >= closest {
					R--
					continue
				} else {
					closest = tmp
					result = nums[i] + nums[L] + nums[R]
					R--
				}
			}
			if nums[i]+nums[L]+nums[R] < target {
				tmp := target - (nums[i] + nums[L] + nums[R])
				//fmt.Println(tmp, target, result)
				if tmp >= closest {
					L++
					continue
				} else {
					closest = tmp
					result = nums[i] + nums[L] + nums[R]
					L++
				}
			}
		}
	}
	return result
}

func TestThreeSumClosest(t *testing.T) {
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
}
