package problem

import (
	"fmt"
	"strconv"
	"testing"
)

func restoreIpAddresses(s string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	return recursiveIp("", s, 4-1)
}

func recursiveIp(pre string, s string, n int) []string {
	var result []string
	length := len(s)
	var TempPre string
	if length > 0 {
		for i := 1; i <= 3; i++ {
			if i > length {
				break
			}
			current, _ := strconv.Atoi(s[length-i : length])
			old, _ := strconv.Atoi(s[length-i+1 : length])
			if current == old && i != 1 {
				continue
			}
			fmt.Println(current)
			if current <= 255 && n <= len(s[:length-i]) && len(s[:length-i]) <= n*3 {
				if pre == "" {
					TempPre = strconv.Itoa(current)
				} else {
					TempPre = strconv.Itoa(current) + "." + pre
				}
				result = append(result, recursiveIp(TempPre, s[:length-i], n-1)...)
			}
		}
	} else {
		result = append(result, pre)
	}
	fmt.Println(result)
	return result
}

func TestRecursiveIp(t *testing.T) {
	fmt.Println(restoreIpAddresses("1111"))
}
