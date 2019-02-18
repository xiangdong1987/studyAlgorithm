package problem

/*
* 求众数
 */
func majorityElement(nums []int) int {
	num := len(nums) / 2
	numMap := make(map[int]int)
	for _, val := range nums {
		numMap[val] = numMap[val] + 1
		if numMap[val] > num {
			return val
		}
	}
	return 0
}

//比较好的算法摩尔投票法（在肯定有众数的前提下摩尔投票法得到的解是众数）
func BoyerMooreVoting(nums []int) int {
	counter := 0
	var now = nums[0]
	for _, val := range nums {
		if counter == 0 {
			now = val
		}
		if now == val {
			counter++
		} else {
			counter--
		}
	}
	return now
}
