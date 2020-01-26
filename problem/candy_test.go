package problem

import (
	"math"
	"testing"
)

func candy(ratings []int) int {
	leftToRight := make([]int, len(ratings))
	RightToLeft := make([]int, len(ratings))
	sum := 0
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			leftToRight[i] = 1
		} else {
			if ratings[i] > ratings[i-1] {
				leftToRight[i] = leftToRight[i-1] + 1
			} else {
				leftToRight[i] = 1
			}
		}
	}

	for j := len(ratings) - 1; j >= 0; j-- {
		if j == len(ratings)-1 {
			RightToLeft[j] = 1
		} else {
			if ratings[j] > ratings[j+1] {
				RightToLeft[j] = RightToLeft[j+1] + 1
			} else {
				RightToLeft[j] = 1
			}
		}
	}

	for k := 0; k < len(leftToRight); k++ {
		sum += int(math.Max(float64(leftToRight[k]), float64(RightToLeft[k])))
	}
	return int(sum)
}
func candy1(ratings []int) int {
	candies := 0
	var status int
	var continueNum int
	for i := 1; i < len(ratings); i++ {
		if i == 1 {
			if ratings[i] > ratings[i-1] {
				status = 1
				continueNum = 2
			} else if ratings[i] < ratings[i-1] {
				status = 2
				continueNum = 2
			} else {
				status = 3
				continueNum = 1
			}
		} else {
			if ratings[i] > ratings[i-1] {
				if status == 1 {
					continueNum++
				} else if status == 2 {
					status = 1
					candies += continueNum*(1+continueNum)/2 - 1
					continueNum = 2
				} else {
					status = 1
					candies += continueNum*(1+continueNum)/2 - 1
					continueNum = 2
				}
			} else if ratings[i] < ratings[i-1] {
				if status == 2 {
					continueNum++
				} else if status == 1 {
					status = 2
					candies += continueNum * (1 + continueNum) / 2
					continueNum = 2
				} else {
					status = 2
					candies += continueNum*(1+continueNum)/2 - 1
					continueNum = 2
				}
			} else {
				if status == 3 {
					candies++
				} else if status == 2 {
					status = 3
					candies += continueNum*(1+continueNum)/2 - 1
					continueNum = 2
				} else {
					status = 3
					candies += continueNum*(1+continueNum)/2 - 1
					continueNum = 2
				}
			}
		}
		if i == len(ratings)-1 {
			if status == 3 {
				candies++
			} else {
				candies += continueNum * (1 + continueNum) / 2
			}
		}
		println(candies)
	}
	return candies
}

func TestCandy(t *testing.T) {
	println(candy([]int{12, 4, 3, 11, 34, 34, 1, 67}))
}
