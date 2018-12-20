package stringMatch

import "fmt"

func TestKmp() {
	a := "ababaeabacaaaaaddfdfdfdfdf"
	b := "aca"
	pos := Kmp(b, a)
	fmt.Println(pos)
}

func Kmp(needle string, str string) int {
	next := getNext(needle)
	fmt.Println(next)
	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != needle[j] {
			j = next[j-1] + 1
		}
		if str[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - j + 1
		}
	}
	return -1
}

func getNext(needle string) []int {
	var next = make([]int, len(needle))
	fmt.Println(next)
	next[0] = -1
	k := -1
	for i := 1; i < len(needle); i++ {
		for k != -1 && needle[k+1] != needle[i] {
			k = next[k]
		}
		if needle[k+1] == needle[i] {
			k++
		}
		next[i] = k
	}
	return next
}
