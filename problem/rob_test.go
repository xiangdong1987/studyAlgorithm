package problem

import "math"

/**
198 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
问题分析：
1.求最大值问题，一般很可能跟动态规划有关
2.我们尝试进行推导动态转移方程式，假设数据元素未n dp来存储当前最大状态
	n=1 直接返回第一个值 dp[0]=num[0]
	n=2 返回两个中较大的 dp[1]=max(num[0],num[1])
	n=3 max(dp[i-2]+num[i],dp[i-1])
	dp[i]=max(dp[i-2]+num[i],dp[i-1])
有了这个转移方程式我们就可以写代码了
*/
func rob(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else if i == 1 {
			dp[i] = int(math.Max(float64(nums[i]), float64(nums[i-1])))
		} else {
			dp[i] = int(math.Max(float64(dp[i-2]+nums[i]), float64(dp[i-1])))
		}
	}
	return dp[len(nums)-1]
}
