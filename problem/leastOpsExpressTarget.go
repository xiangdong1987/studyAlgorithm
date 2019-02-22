package problem

import (
	"fmt"
	"math"
)

/**
分析：
1.乘法是使用最少能达到最目标值的运算，所以应该从乘法开始计算
2.以使用运算符多少为权重进行计算
3.权重分析 乘法为x^n n为权重 x+x 2 > x/x 2 > x-x 2
4.递归的求最小的权重
*/
func TestLeastOpsExpressTarget() {
	n := leastOpsExpressTarget(3, 365)
	fmt.Println(n)
}
func leastOpsExpressTarget(x int, target int) int {

	if x > target {
		return int(math.Min(float64(2*target-1), float64((x-target)*2)))
	} else {

		cost := 0
		t := x
		for x*t < target {
			cost += 1
			t *= x
		}
		cost++
		fmt.Println(x, target, t)
		if t*x == target {
			return cost
		}
		if t*x-target < target {
			return int(math.Min(float64(cost+leastOpsExpressTarget(x, target-t)), float64(cost+1+leastOpsExpressTarget(x, t*x-target))))
		}
		return cost + leastOpsExpressTarget(x, target-t)
	}
}
