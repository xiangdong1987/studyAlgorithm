package problem

import (
	"fmt"
	"testing"
)

func permute(nums []int) [][]int {
	var pre []int
	return recursivePremute(pre, nums)
}

func recursivePremute(pre []int, left []int) [][]int {
	if len(left) > 0 {
		var result [][]int
		xLeft := left
		for i := 0; i < len(left); i++ {
			tempPre := append(pre, xLeft[i])
			var tempLeft []int
			for j := 0; j < len(left); j++ {
				if left[j] != left[i] {
					tempLeft = append(tempLeft, left[j])
				}
			}
			fmt.Println(tempPre, tempLeft, xLeft, i)
			if len(tempLeft) == 0 {
				result = append(result, tempPre)
				break
			} else {
				result = append(result, recursivePremute(tempPre, tempLeft)...)
			}
		}
		return result
	}
	return [][]int{}
}

func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}
