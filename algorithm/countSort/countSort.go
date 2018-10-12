package CS

import (
	"../../tool"
	"../bubbleSort"
	"fmt"
)

func CountSort() {
	arrList := []int{1, 2, 8, 3, 6, 8, 4, 3, 3}
	arrList = standarCountSort(arrList)
	fmt.Println("CountSort:")
	tool.PrintSlice(arrList)
}

func standarCountSort(list []int) []int {
	return createCountList(list)
}
func createCountList(list []int) []int {
	var countList map[int]int
	countList = make(map[int]int)
	i := 0
	for i < len(list) {
		if _, ok := countList[list[i]]; ok {
			countList[list[i]]++
		} else {
			countList[list[i]] = 1
		}
		i++
	}
	countSorted := []int{}
	for key := range countList {
		//fmt.Println(key, "key:", countList[key])
		countSorted = append(countSorted, key)
	}
	BS.StandardBubbleSort(countSorted)
	tool.PrintSlice(countSorted)
	var countListSorted map[int]int
	countListSorted = make(map[int]int)
	i = 0
	sum := 0
	for i < len(countSorted) {
		sum = sum + countList[countSorted[i]]
		countListSorted[countSorted[i]] = sum
		i++
	}
	//反向查询
	resultList := make([]int, 9)
	i = 0
	for i < len(list) {
		countListSorted[list[i]]--
		resultList[countListSorted[list[i]]] = list[i]
		i++
	}
	return resultList
}
