package problem

import (
	"fmt"
	"testing"
)

func climbStairs1(n int) int {
	sum := 0
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	sum += climbStairs(n - 1)
	sum += climbStairs(n - 2)
	return sum
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
func TestClimbStairs(t *testing.T) {
	fmt.Println(climbStairs(44))
}
