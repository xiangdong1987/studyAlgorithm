package baseMind

import "fmt"

func TestDynOneZero() {
	max := dynOneZero()
	fmt.Println(max)
}
func dynOneZero() int {
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
	fmt.Println(statMap)
	for m := bagWeight; m > 0; m-- {
		if statMap[itemNum-1][m] {
			return m
		}
	}

	return 0
}
