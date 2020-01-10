package problem

import (
	"fmt"
	"testing"
)

const INT32_MAX = int(^uint32(0) >> 1)
const INT32_MIN = ^INT32_MAX

func reverse(x int) int {
	result := 0
	for x != 0 {
		if result*10 < int(INT32_MAX) && result*10 > INT32_MIN {
			current := x % 10
			result = result*10 + current
			x = x / 10
		} else {
			return 0
		}
	}
	return result
}

func TestReverse(t *testing.T) {
	fmt.Print(reverse(123))
}
