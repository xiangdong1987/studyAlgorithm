package problem

import (
	"fmt"
	"math"
	"testing"
)

/**
55 给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
问题分析：
1.分析题目要求，因为要不停去查找可到达情况，所以需要回溯和递归，动态的进行规划
2.动态规划问题大部分利用的是动态转移或者递归回溯，题目是多种结果的一种
用转移的方无法确定是哪一种，说以应该用递归回溯的方式进行检查
3.利用递归回溯方式的动态规划主要有下面的优化思路
	普通回溯：首先，用原始的递归回溯方式，遍历所有可能性递归，这种方式思路便于理解，
	记忆回溯：但是大量数据的时候，可能性过多造成程序执行时间过长，为了优化我们可以采取增加额外
	空间保存点的可达到情况，避免重复计算优化执行速度
	自底向上：上面两种情况都是自顶向下的递归，在使用记忆回溯的时候需要双倍的额外空间来记录执行情况
	在此题中，我们分析可以得到，如果我们从右侧想左测递归的话我们可以避免回溯，因为我们右侧都是计算
	好的，无需在返回再计算
	贪心算法：我们从左向右找最大的可触及距离，如果最远距离小于小于当前的距离就说明不可触达到最后。大家可以仔细理解一下。
*/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	max := 0
	for i := 1; i <= nums[0]; i++ {
		if nums[i]+i < max {
			continue
		} else {
			max = nums[i] + i
		}
		res := recursiveCanJump(nums, i)
		if res {
			return res
		}
	}
	return false
}
func recursiveCanJump(nums []int, pos int) bool {
	current := nums[pos]
	length := len(nums)
	fmt.Println(pos, current, length)
	if pos+current+1 >= length {
		return true
	} else {
		for i := pos + 1; i < pos+current+1; i++ {
			res := recursiveCanJump(nums, i)
			if res {
				return true
			}
		}
		return false
	}
}

//贪婪算法
func canJump1(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		max = int(math.Max(float64(max), float64(nums[i]+i)))

	}
	return true
}

func TestCanJump(t *testing.T) {
	fmt.Println(canJump([]int{1, 2, 3}))
}
