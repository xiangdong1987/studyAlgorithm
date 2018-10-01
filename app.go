package main

import (
	"./insertSort"
	"./quickSort"
)

func main() {
	//原始快排
	Qs.OriginalQuickSort()
	//优化快排
	Qs.OptimizationQuickSort()
	//插入算法
	IS.InsertSort()
}
