package IS

import "fmt"
import "../tool"

func InsertSort() {
	arrList := []int{1, 2, 3, 4, 5, 6, 8, 9, 343, 3}
	standardInsertSort(arrList)
	fmt.Println("InsertSort:")
	tool.PrintSlice(arrList)
}

func standardInsertSort(list []int) {
	length := len(list)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if list[i] <= list[j] {
				tool.SliceSwap(list, i, j)
			}
		}
	}
}
