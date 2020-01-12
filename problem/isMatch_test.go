package problem

import (
	"fmt"
	"testing"
)

//字符创正则
func isMatch(s string, p string) bool {
	//当前字符串是否匹配
	firstMatch := false
	if len(p) < 1 {
		return len(s) < 1
	}
	if len(s) > 0 && (s[0] == p[0] || p[:1] == ".") {
		firstMatch = true
	}
	if len(p) >= 2 && p[1:2] == "*" {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func dp(i int, j int, s string, p string) {

}

func TestIsMatch(t *testing.T) {
	fmt.Print(isMatch("aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"))
}
