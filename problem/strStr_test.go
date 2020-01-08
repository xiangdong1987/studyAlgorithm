package problem

import "testing"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if i+j > len(haystack) {
				break
			}
			if haystack[i+j] == needle[j] {
				if j == len(needle)-1 {
					return i
				}
				continue
			} else {
				break
			}
		}
	}
	return -1
}
func TestStrStr(t *testing.T) {
	println(strStr("hello", "ll"))
}
