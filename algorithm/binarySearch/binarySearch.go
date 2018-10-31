package binarySearch

import "fmt"

func BinarySearch() {
	list := []int{3, 4, 7, 9, 10, 12, 15}
	result := binarySearch(list, 1)
	fmt.Println("result:", result)
}

func binarySearch(list []int, x int) int {
	length := len(list)
	fmt.Println(list)
	if length == 1 {
		return 0
	}
	min := length / 2
	if list[min] > x {
		return binarySearch(list[0:min], x)
	} else if list[min] < x {
		return binarySearch(list[min:length], x)
	} else {
		return list[min]
	}
}
