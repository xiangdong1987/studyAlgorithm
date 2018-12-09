package BS

import (
	"fmt"
	"studyAlgorithm/tool"
)

func BubbleSort() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 4, 9, 343, 3}
	StandardBubbleSort(arrList)
	fmt.Println("BubbleSort:")
	tool.PrintSlice(arrList)
}
func StandardBubbleSort(list []int) {
	i := len(list)
	for i > 0 {
		for j := 0; j < i-1; j++ {
			if list[j] > list[j+1] {
				tool.SliceSwap(list, j, j+1)
			}
		}
		i--
	}
}
