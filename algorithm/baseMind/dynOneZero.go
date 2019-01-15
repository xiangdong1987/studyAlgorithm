package baseMind

import "fmt"

func TestDynOneZero() {
	max := knapsack()
	fmt.Println("max:", max)
	max = knapsack1()
	fmt.Println("max1:", max)
	max = knapsack2()
	fmt.Println("max2:", max)
}
func knapsack() int {
	var itemNum = 5
	var bagWeight = 10
	var items = []int{4, 5, 2, 7, 1}
	statMap := map[int]map[int]bool{}
	var tmpMap map[int]bool
	map0 := make(map[int]bool, bagWeight)
	map0[0] = true
	statMap[0] = map0
	for i := 1; i < itemNum; i++ {
		tmpMap = make(map[int]bool, bagWeight)
		//不把第i个物品放入背包
		for j := 0; j <= bagWeight; j++ {
			if statMap[i-1][j] {
				tmpMap[j] = statMap[i-1][j]
			}
		}
		//把第i个物品放入背包
		for j := 0; j <= bagWeight-items[i]; j++ {
			//不把第i个物品放入背包
			if statMap[i-1][j] {
				tmpMap[j+items[i]] = statMap[i-1][j]
			}
		}
		statMap[i] = tmpMap
	}
	//fmt.Println(statMap)
	for m := bagWeight; m > 0; m-- {
		if statMap[itemNum-1][m] {
			return m
		}
	}
	return 0
}

func knapsack1() int {
	var itemNum = 5
	var bagWeight = 10
	var items = []int{3, 5, 2, 7, 1}
	statMap := make(map[int]bool, bagWeight)
	statMap[0] = true
	for i := 1; i < itemNum; i++ {
		for j := bagWeight - items[i]; j >= 0; j-- {
			if statMap[j] == true {
				statMap[j+items[i]] = true
			}
		}
	}
	for n := bagWeight; n >= 0; n-- {
		if statMap[n] {
			return n
		}
	}
	return 0
}

func knapsack2() int {
	var itemNum = 5
	var bagWeight = 10
	var items = []int{3, 5, 2, 7, 1}
	var values = []int{10, 2, 15, 20, 13}
	statMap := map[int]map[int]int{}
	var tmpMap map[int]int
	map0 := make(map[int]int, bagWeight)
	map0[0] = values[0]
	statMap[0] = map0
	for i := 1; i < itemNum; i++ {
		tmpMap = make(map[int]int, bagWeight)
		//不选择第i个物品
		for j := 0; j < bagWeight; j++ {
			if statMap[i-1][j] > 0 {
				tmpMap[j] = statMap[i-1][j]
			}
		}
		//选择第i个物品
		for j := 0; j < bagWeight-items[i]; j++ {
			if statMap[i-1][j] > 0 {
				v := statMap[i-1][j] + values[i]
				if v > statMap[i-1][j+items[i]] {
					tmpMap[j+items[i]] = v
				}
			}
		}
		statMap[i] = tmpMap
	}
	fmt.Println(statMap)
	for n := bagWeight; n >= 0; n-- {
		if statMap[itemNum-1][n] > 0 {
			return statMap[itemNum-1][n]
		}
	}
	return 0
}
