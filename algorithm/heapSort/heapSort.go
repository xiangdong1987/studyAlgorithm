package HS

import (
	"../../tool"
	"fmt"
)

func HeapSort() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 4, 9, 343, 3}
	standarHeapSort(arrList)
	fmt.Println("HeapSort:")
	tool.PrintSlice(arrList)
}

func standarHeapSort(list []int) {
	creatHeap(list, len(list))
}
func adjustHeap(list []int, length int) {
	min := length / 2
	//首位交换
	tool.SliceSwap(list, min, length-1)
	creatHeap(list, length-1)
}
func creatHeap(list []int, length int) {
	if length == 2 {
		if list[1] > list[0] {
			tool.SliceSwap(list, 1, 0)
		}
		return
	}
	i := 0
	min := length / 2
	for min-i >= 0 || min+i <= length {
		if min-i >= 0 && list[min-i] < list[min] {
			tool.SliceSwap(list, min-i, min)
		}
		if min+i <= length && list[min+i-1] < list[min] {
			tool.SliceSwap(list, min+1, min)
		}
		i++
	}
	tool.PrintSlice(list)
	adjustHeap(list, length)
}
