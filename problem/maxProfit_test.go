package problem

import (
	"fmt"
	"math"
	"testing"
)

func maxProfit(prices []int) int {
	maxProfit := 0
	length := len(prices)
	if length < 2 {
		return 0
	}
	minVal := prices[0]
	for i := 1; i < len(prices); i++ {
		maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-minVal)))
		minVal = int(math.Min(float64(prices[i]), float64(minVal)))
	}
	return maxProfit
}
func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
