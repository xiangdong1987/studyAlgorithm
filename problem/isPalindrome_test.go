package problem

import (
	"fmt"
	"testing"
)

/**
*回文数
 */
func isPalindrome(x int) bool {
	list := make([]int, 0)

	if x < 0 {
		return false
	}
	for x > 0 {
		current := x % 10
		list = append(list, current)
		x = x / 10
	}
	end := len(list) - 1
	start := 0
	for start < end {
		fmt.Println(start, end)
		if list[start] != list[end] {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	fmt.Print(isPalindrome(10))
}
