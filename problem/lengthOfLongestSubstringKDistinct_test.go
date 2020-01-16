package problem

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	charMap := make(map[int32]bool)
	dp := make(map[int]int)
	dp[0] = 1
	for i, val := range s {
		if i == 0 {
			charMap[val] = true
			continue
		}
		_, ok := charMap[val]
		if ok {

		}
	}
	return 0
}

type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func TestA(t *testing.T) {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
