package problem

import (
	"fmt"
	"sort"
	"testing"
)

/**
39 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为
target 的组合。candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
问题分析：
1. 看到要和等于目标数，这样子就要枚举进行匹配，并且解不能重复，那就是出现过的不能再出现了
2. 第一个想到的是递归，没次如果当前值小于等于放入组合中就将目标值减小，再次尝试递归，
   最后等于的时候加入解集
3. 开始组的时候发现他会回头，因为我们传入的是整个数组，我们为了避免回头，我们首先要排序
   然后每次都放入当前节点及以后的数据，这样就避免了回头
总结起来：其实这就是经典的递归回溯问题，每次深度尝试，并且去掉不合适的解，循环递归
我总结为：不撞南墙不回头算法

*/
var total [][]int

func combinationSum(candidates []int, target int) [][]int {
	total = make([][]int, 0)
	sort.Ints(candidates)
	recursiveCombinationSum(candidates, target, []int{})
	return total
}

func recursiveCombinationSum(candiates []int, target int, preList []int) {
	fmt.Println(preList, target)
	for i := 0; i < len(candiates); i++ {
		if candiates[i] < target {
			var newList []int
			if len(preList) == 0 {
				newList = []int{candiates[i]}
			} else {
				newList = make([]int, len(preList))
				copy(newList, preList)
				fmt.Println(newList)
				newList = append(newList, candiates[i])
			}
			recursiveCombinationSum(candiates[i:], target-candiates[i], newList)
			//recursiveCombinationSum(candiates[i+1:], target-candiates[i], newList)
		} else if candiates[i] == target {
			now := append(preList, candiates[i])
			total = append(total, now)
			return
		} else {
			return
		}
	}
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
