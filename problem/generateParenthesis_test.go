package problem

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	result = recursiveGenerateParenthesis("", 0, 0, n)
	fmt.Println(result)
	return result

}
func recursiveGenerateParenthesis(str string, open int, close int, max int) []string {
	var result []string
	if len(str) == max*2 {
		fmt.Println(str)
		result = append(result, str)
		return result
	}
	if open < max {
		result = append(result, recursiveGenerateParenthesis(str+"(", open+1, close, max)...)
	}
	if close < open {
		result = append(result, recursiveGenerateParenthesis(str+")", open, close+1, max)...)
	}
	return result
}

func TestGenerate(t *testing.T) {
	fmt.Println(generateParenthesis(3))
}
