package RS

import (
	"../../tool"
	"fmt"
)

type bucket map[int][]int

//利用10进制做一百以内的数排序
func RadixSort() {
	arrList := []int{1, 51, 67, 75, 91, 82, 31, 33, 41, 20}
	resultArray := make([]int, 10)
	resultList := make(map[int]int, len(arrList))
	sortByPosition(arrList, 1, resultList)
	sortByPosition(arrList, 2, resultList)
	for key := range resultList {
		resultArray[resultList[key]] = key
		//fmt.Println(key, "rank is ", resultList[key])
	}
	fmt.Println("RadixSort:")
	tool.PrintSlice(resultArray)
}

func sortByPosition(list []int, postion int, resultList map[int]int) {
	bucketList := make(bucket)
	oldResult := make(map[int]int)
	oldResult = resultList
	radix := 0
	if postion == 1 {
		radix = tool.Power(10, postion)
	} else {
		radix = tool.Power(10, postion-1)
	}
	//初始化桶
	for j := 0; j < 10; j++ {
		bucketList[j] = []int{}
	}
	x := 0
	for i := 0; i < len(list); i++ {
		if postion == 1 {
			x = list[i] % radix
		} else {
			x = list[i] / radix
		}
		//fmt.Println(x)
		bucketList[x] = append(bucketList[x], list[i])
		//tool.PrintSlice(bucketList[x])
	}
	//桶根据原始排名排序
	if postion > 1 {
		for j := 0; j < 10; j++ {
			bucketSort(bucketList[j], oldResult)
		}
	}
	rank := 0
	for j := 0; j < 10; j++ {
		for n := 0; n < len(bucketList[j]); n++ {
			resultList[bucketList[j][n]] = rank
			rank++
		}
	}
}
func bucketSort(list []int, resultMap map[int]int) {
	for j := len(list); j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if resultMap[list[i+1]] < resultMap[list[i]] {
				tool.SliceSwap(list, i, i+1)
			}
		}
	}
}
