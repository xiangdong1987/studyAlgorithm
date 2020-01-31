package problem

import (
	"fmt"
	"sort"
	"testing"
)

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)
	for i := 0; i < len(word2)+1; i++ {
		dp[i] = make([]int, len(word1)+1)
	}
	for i := 1; i < len(word1)+1; i++ {
		dp[0][i] = dp[0][i-1] + 1
	}
	for j := 1; j < len(word2)+1; j++ {
		dp[j][0] = dp[j-1][0] + 1
	}
	fmt.Println(dp)
	for n := 1; n <= len(word2); n++ {
		for m := 1; m <= len(word1); m++ {
			if word2[n-1] == word1[m-1] {
				dp[n][m] = dp[n-1][m-1]
			} else {
				temp := []int{dp[n-1][m], dp[n-1][m-1], dp[n][m-1]}
				sort.Ints(temp)
				dp[n][m] = temp[0] + 1
			}
		}
	}
	fmt.Println(dp)
	return dp[len(word2)][len(word1)]
}
func TestMinDistance(t *testing.T) {
	fmt.Println(minDistance("horse", "ros"))
}
