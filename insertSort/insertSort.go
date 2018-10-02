package IS

import "fmt"
import "../tool"

func InsertSort() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 9, 343, 3}
	standardInsertSort(arrList)
	fmt.Println("InsertSort:")
	tool.PrintSlice(arrList)
}

func standardInsertSort(list []int) {
	length := len(list)
	i := 0
	for i < length {
		tmp := list[i]
		for j := i + 1; j < length; j++ {
			if tmp <= list[j] {
				tool.SliceSwap(list, i, j)
				tool.SliceSwap(list, i+1, j)
				i++
			}

		}
	}
}
