package Qs

import (
	"fmt"
)
import "../../tool"

func OriginalQuickSort() {
	arrList := []int{1, 2, 3, 4, 5, 6, 8, 9, 343, 3}
	standardQuickSort(arrList, 0, len(arrList))
	fmt.Println("OriginalQuickSort:")
	tool.PrintSlice(arrList)

}
func OptimizationQuickSort() {
	arrList := []int{1, 2, 3, 4, 5, 6, 8, 9, 343, 3}
	randQuickSort(arrList, 0, len(arrList))
	fmt.Println("OptimizationQuickSort:")
	tool.PrintSlice(arrList)
}

func standardQuickSort(list []int, left int, right int) {
	if left >= right {
		return
	}
	temp := list[left]
	oldLeft := left
	newLeft := left
	oldRight := right
	j := right
	for i := left + 1; i < j; i++ {
		if list[i] < temp {
			newLeft++
			tool.SliceSwap(list, newLeft-1, i)
		}
	}
	tool.PrintSlice(list)
	//判断已经是最小值了
	if newLeft == oldLeft {
		standardQuickSort(list, oldLeft+1, oldRight)
	} else {
		standardQuickSort(list, oldLeft, newLeft)
		standardQuickSort(list, newLeft, oldRight)
	}
}
func randQuickSort(list []int, left int, right int) {
	if left >= right {
		return
	}
	newLeft := tool.GenerateRangeNum(left, right)
	temp := list[newLeft]
	oldLeft := left
	oldRight := right
	j := right
	for i := left; i < j; i++ {
		if i < newLeft {
			if list[i] > temp {
				newLeft = i
				tool.SliceSwap(list, newLeft, i)
			}
		} else if i == newLeft {
		} else {
			if list[i] <= temp {
				tool.SliceSwap(list, newLeft, i)
				tool.SliceSwap(list, newLeft+1, i)
				newLeft++
			}
		}
	}
	tool.PrintSlice(list)
	//判断已经是最小值了
	if newLeft == oldLeft {
		randQuickSort(list, oldLeft+1, oldRight)
	} else {
		randQuickSort(list, oldLeft, newLeft)
		randQuickSort(list, newLeft, oldRight)
	}
}
func RecursiveQuickSort() {
	arrList := []int{1, 2, 6, 8, 9, 3, 4, 5, 343}
	arrList = divisionData(arrList)
	fmt.Println("OptimizationQuickSort:")
	tool.PrintSlice(arrList)
}

func divisionData(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	var left []int
	var right []int
	mid := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] <= mid {
			left = append(left, list[i])
		} else {
			right = append(right, list[i])
		}
	}
	left = append(left, mid)
	left = divisionData(left)
	right = divisionData(right)
	return append(left, right...)
}
