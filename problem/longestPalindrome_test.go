package problem

import (
	"fmt"
	"math"
	"testing"
)

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
func TestLongestPalindrome(t *testing.T) {
	n := longestPalindrome("babad")
	fmt.Println(n)
}
