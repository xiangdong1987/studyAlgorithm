package baseMind

import (
	"fmt"
)

var totalNum int

func CountReversed() {
	list := []int{1, 2, 43, 344, 3, 46, 75}
	//list := []int{5, 2, 43, 344, 3, 46, 75}
	mergeSortCount(list, 0, len(list)-1)
	fmt.Println(totalNum)
}
func mergeSortCount(a []int, start int, end int) {
	//return
	if start >= end {
		return
	}
	low := start
	min := (start + end) / 2
	mergeSortCount(a, low, min)
	mergeSortCount(a, min+1, end)
	mergeCount(a, low, min, end)
}
func mergeCount(a []int, low int, min int, end int) {
	tmp := make([]int, end-low+1)
	i := low
	k := 0
	j := min + 1
	//获取左半边和右半边的逆序度
	for i <= min && j <= end {
		if a[i] < a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			totalNum += min - i + 1
			tmp[k] = a[j]
			k++
			j++
		}
	}
	//处理左边剩下
	//fmt.Println(i, min)
	for i <= min {
		tmp[k] = a[i]
		k++
		i++
	}
	//处理右边剩下
	//fmt.Println(j, end)
	for j <= end {
		tmp[k] = a[j]
		k++
		j++
	}
	fmt.Println(tmp, low, min, end)
	//将数据拷贝回a
	for n := 0; n < end-low+1; n++ {
		//fmt.Println(n)
		a[low+n] = tmp[n]
	}
	fmt.Println(a)
	fmt.Println("num:", totalNum)
}
