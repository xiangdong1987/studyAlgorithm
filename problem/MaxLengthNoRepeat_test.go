package problem

import "testing"

//暴力解决
func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	hadChar := make(map[uint8]uint8, len(s))
	for i := 0; i < len(s); i++ {
		if len(hadChar) > maxLength {
			maxLength = len(hadChar)
		}
		hadChar = make(map[uint8]uint8, len(s))
		for j := i; j < len(s); j++ {
			_, ok := hadChar[s[j]]
			if ok {
				//清空map
				break
			} else {
				hadChar[s[j]] = s[j]
			}
		}
	}
	if len(hadChar) > maxLength {
		maxLength = len(hadChar)
	}
	return maxLength
}

//滑动窗口
func lengthOfLongestSubstringB(s string) int {
	var maxLength = 0
	sliceWin := make(map[uint8]uint8, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		for j < len(s) {
			_, ok := sliceWin[s[j]]
			if ok {
				if len(sliceWin) > maxLength {
					maxLength = len(sliceWin)
				}
				//清空map
				delete(sliceWin, s[i])
				break
			} else {
				sliceWin[s[j]] = s[j]
				j++
			}
		}
	}
	if len(sliceWin) > maxLength {
		maxLength = len(sliceWin)
	}
	return maxLength
}
func TestMaxLength(t *testing.T) {
	str := "jbpnbwwd"
	println(lengthOfLongestSubstring(str))
	println(lengthOfLongestSubstringB(str))
}
