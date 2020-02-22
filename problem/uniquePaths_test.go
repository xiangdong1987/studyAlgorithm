package problem

import (
	"fmt"
	"testing"
)

/**
67 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
问题分析：
1.现在不管拿到什么题我都会先想用递归来尝试，果然递归可以并且代码也异常的简单
2.但是问题来了，当网格边大时，可能性成指数增长，当一个问题变成指数级解，那就不可能是最好的方式了
3.这时候就开始考虑高级的动态规划，怎么下手呢？动态规划就是问题拆分，递归是从先往后拆分不成立
4.那我们就从后往前拆分看看，我们如果假设dp[i][j]代表dp[i][j]位最多可能数，那它等于什么呢？
5.因为我们只能往右或者往下走所以 dp[i][j]=dp[i-1][j]+dp[i][j-1]
6.一下子我们的世界都明亮了，我们只要知道dp[i-1][j]和dp[i][j-1]就可以了
7.我们的上边界 dp[0][0-m]=1 左边界 dp[0-n][0]=1
8.最后我们可以连续推导出所有的情况，并且不会重复计算使得算法得到优化
总结：思路变通很重要，正向不行就逆向 只要不走重复路我们就是最棒的
*/

func uniquePathsNew(m int, n int) int {
	dp := make([][]int, n)
	//初始化动态存储
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = 1
	//边界问题
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		fmt.Println(dp)
		for j := 1; j < m; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[n-1][m-1]
}

var sum int

func uniquePaths(m int, n int) int {
	recursiveUniquePaths(m-1, n-1)
	return sum
}

func recursiveUniquePaths(m int, n int) {
	fmt.Println(m, n)
	if m > 0 {
		recursiveUniquePaths(m-1, n)
	}
	if n > 0 {
		recursiveUniquePaths(m, n-1)
	}
	if m == 0 && n == 0 {
		sum++
		return
	}
}

func TestUniquePaths(t *testing.T) {
	fmt.Println(uniquePaths(7, 3))
}
