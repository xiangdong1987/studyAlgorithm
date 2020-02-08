package problem

import (
	"fmt"
	"testing"
)

func longestValidParentheses(s string) int {
	open := 0
	close := 0
	max := 0
	for i := 0; i < len(s); i++ {
		open = 0
		close = 0
		currentMax := 0
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				open++
			} else {
				close++
			}
			if j == len(s)-1 {
				if open != close {
					break
				}
			}
			if open >= close {
				currentMax++
				if open == close {
					if currentMax > max {
						max = currentMax
					}
				}
			} else {
				break
			}
		}
	}
	return max
}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses(")()())"))
}
