package problem

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numberMap := make(map[uint8][]string, 8)
	numberMap['2'] = []string{"a", "b", "c"}
	numberMap['3'] = []string{"d", "e", "f"}
	numberMap['4'] = []string{"g", "h", "i"}
	numberMap['5'] = []string{"j", "k", "l"}
	numberMap['6'] = []string{"m", "n", "o"}
	numberMap['7'] = []string{"p", "q", "r", "s"}
	numberMap['8'] = []string{"t", "u", "v"}
	numberMap['9'] = []string{"w", "x", "y", "z"}
	if len(digits) == 1 {
		return numberMap[digits[0]]
	}
	var leftString string
	if len(digits) == 2 {
		leftString = ""
	} else {
		leftString = digits[2:]
	}
	return recursiveLetterCombinations(numberMap, numberMap[digits[0]], digits[1], leftString)
}

func recursiveLetterCombinations(numberMap map[uint8][]string, preList []string, current uint8, leftString string) []string {
	var all []string
	list, ok := numberMap[current]
	if !ok {
		return []string{}
	}
	newPre := []string{}
	for j := 0; j < len(preList); j++ {
		for i := 0; i < len(list); i++ {
			newPre = append(newPre, preList[j]+list[i])
		}
	}
	if len(leftString) > 0 {
		all = append(all, recursiveLetterCombinations(numberMap, newPre, leftString[0], leftString[1:])...)
	} else {
		all = preList
	}
	return all
}

func TestLetterCombination(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}
