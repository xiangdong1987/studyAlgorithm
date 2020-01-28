package problem

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	stack := make([]int, 0)
	left := map[string]int{
		"{": 1,
		"(": 2,
		"[": 3,
	}
	right := map[string]int{
		"}": 1,
		")": 2,
		"]": 3,
	}
	for _, value := range s {
		fmt.Println(stack)
		lIndex, ok := left[string(value)]
		if ok {
			stack = append(stack, lIndex)
		} else {
			rIndex, ok := right[string(value)]
			if ok {
				if len(stack) == 0 {
					return false
				}
				leftIndex := stack[len(stack)-1]
				if leftIndex == rIndex {
					stack = append(stack[:len(stack)-1], stack[len(stack):]...)
				} else {
					return false
				}
			}
		}
	}
	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("]"))
}
