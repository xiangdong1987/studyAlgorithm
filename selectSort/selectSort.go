package SS

import (
	"fmt"
)
import (
	"../tool"
)

func SelectSort() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 4, 9, 343, 3}
	arrList = standardSelectSort(arrList)
	fmt.Println("SelectSort:")
	tool.PrintSlice(arrList)
}

func standardSelectSort(arrList []int) []int {
	i := 0
	for i < len(arrList) {
		mini := arrList[i]
		for j := i + 1; j < len(arrList); j++ {
			if arrList[j] <= mini {
				tool.SliceSwap(arrList, i, j)
			}
		}
		i++
	}
	return arrList
}
