package problem

import (
	"fmt"
	"sort"
	"testing"
)

//贪婪算法会漏掉一些结果
//需要用动态规划，因为有局部最优解
//可以更具变换方程式获得
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	notebook := make([]int, amount)
	return recursiveCoinChange(coins, amount, notebook)

}
func recursiveCoinChange(coins []int, amount int, notebook []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	min := INT32_MAX
	if notebook[amount-1] != 0 {
		return notebook[amount-1]
	}
	for i := len(coins) - 1; i >= 0; i-- {
		res := recursiveCoinChange(coins, amount-coins[i], notebook)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == INT32_MAX {
		notebook[amount-1] = -1
	} else {
		notebook[amount-1] = min
	}
	return notebook[amount-1]
}
func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
