package problem

import "math"

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfitNew(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	//动态规划
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = 0 - prices[0]
	//0 代表不买 1代表买
	for j := 1; j < len(prices); j++ {
		dp[j][0] = int(math.Max(float64(dp[j-1][0]), float64(dp[j-1][1]+prices[j])))
		dp[j][1] = int(math.Max(float64(dp[j-1][1]), float64(dp[j-1][0]-prices[j])))
	}
	return dp[len(prices)-1][0]
}
