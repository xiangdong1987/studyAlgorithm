package problem

import (
	"fmt"
	"testing"
)

func maxArea(height []int) int {
	max := 0
	for i := 1; i < len(height); i++ {
		for j := i - 1; j >= 0; j-- {
			fmt.Println(max)
			if height[i] > height[j] {
				if max < height[j]*(i-j) {
					max = height[j] * (i - j)
				}
			} else {
				if max < height[i]*(i-j) {
					max = height[i] * (i - j)
				}
			}
		}
	}
	return max
}
func TestMaxArea(t *testing.T) {
	n := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(n)
}
