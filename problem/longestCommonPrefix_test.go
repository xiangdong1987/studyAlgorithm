package problem

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			if len(prefix) > 0 {
				prefix = prefix[:len(prefix)-1]
			} else {
				return ""
			}
		}
	}
	return prefix
}
