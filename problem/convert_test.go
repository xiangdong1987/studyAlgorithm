package problem

import (
	"fmt"
	"testing"
)

/**
z字形转换
*/
func convert(s string, numRows int) string {
	current := 0
	result := make([]uint8, len(s))
	strLength := len(s)
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 {
			pos := i
			for pos < strLength && current < strLength {
				result[current] = s[pos]
				current++
				pos = pos + numRows + numRows - 2
			}
		} else {
			first := i
			for first < strLength && current < strLength {
				fmt.Println(first)
				result[current] = s[first]
				current++
				if current == strLength {
					break
				}
				if first+2*numRows-2-2*i < strLength {
					result[current] = s[first+2*numRows-2-2*i]
					current++
					first = first + numRows + numRows - 2
				} else {
					break
				}
			}
		}
	}
	return string(result)
}

func TestConvert(t *testing.T) {
	fmt.Print(convert("LEETCODEISHIRING", 4))
}
