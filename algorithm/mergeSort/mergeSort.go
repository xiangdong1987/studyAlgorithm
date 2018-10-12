package MS

import (
	"../../tool"
	"fmt"
)

func MergeSort() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 4, 9, 343, 3}
	arrList = standardMergeSort(arrList)
	fmt.Println("MergeSort:")
	tool.PrintSlice(arrList)
}
func standardMergeSort(list []int) []int {
	return divideList(list)
}

func divideList(list []int) []int {
	mid := len(list) / 2
	if mid > 1 {
		left := []int{}
		right := []int{}
		left = divideList(list[:mid])
		right = divideList(list[mid:])
		return mergeList(left, right)
	}
	return list
}

func mergeList(left []int, right []int) []int {
	newList := []int{}
	j := 0
	i := len(left)
	for i > 0 {
		for j = 0; j < len(left)-1; j++ {
			if left[j] >= left[j+1] {
				tool.SliceSwap(left, j, j+1)
			}
		}
		i--
	}

	i = len(right)
	j = 0
	for i > 0 {
		for j := 0; j < len(right)-1; j++ {
			if right[j] >= right[j+1] {
				tool.SliceSwap(right, j, j+1)
			}
		}
		i--
	}
	pa := 0
	pb := 0
	tool.PrintSlice(left)
	tool.PrintSlice(right)
	//return nil
	for pa < len(left) {
		if pb == len(left) {
			if pa < len(left) {
				newList = append(newList, left[pa])
				pa++
				continue
			}
		}
		for pb < len(right) {
			if pa == len(left) {
				if pb < len(right) {
					newList = append(newList, right[pb])
					pb++
					continue
				}
			}
			if left[pa] < right[pb] {
				newList = append(newList, left[pa])
				pa++
			} else if left[pa] == right[pb] {
				newList = append(newList, left[pa])
				newList = append(newList, right[pb])
				pa++
				pb++
			} else {
				newList = append(newList, right[pb])
				pb++
			}
		}
	}
	tool.PrintSlice(newList)
	return newList
}
