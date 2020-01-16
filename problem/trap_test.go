package problem

import (
	"fmt"
	"testing"
)

func trap(height []int) int {
	i := 0
	j := len(height) - 1
	h := 1
	result := 0
	for i < j {
		for i < len(height) && height[i] < h {
			i++
		}
		for j > 0 && height[j] < h {
			j--
		}
		k := i + 1
		for k < j {
			if height[k] >= h {
				k++
				continue
			} else {
				result++
				k++
			}
		}
		h++
	}
	return result
}

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
