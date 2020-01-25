package problem

import (
	"fmt"
	"math"
	"testing"
)

func findKthNumber(n int, k int) int {
	currentNum := 1
	prefix := 1
	for currentNum < k {
		count := getCount(prefix, n)
		if currentNum+count > k {
			prefix *= 10
			currentNum++
		} else {
			prefix++
			currentNum += count
		}
	}
	return prefix
}
func getCount(prefix int, n int) int {
	count := 0
	cur := prefix
	next := prefix + 1
	for cur <= n {
		count += int(math.Min(float64(next), float64(n+1))) - cur
		cur *= 10
		next *= 10
	}
	return count
}
func TestFindKthNumber(t *testing.T) {
	fmt.Println(findKthNumber(10000, 10000))
}
