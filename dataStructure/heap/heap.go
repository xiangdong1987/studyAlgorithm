package heap

import (
	"fmt"
)

type Node struct {
	Value int
	Key   string
}

func CreateHeap() {

	arrList := []int{1, 2, 11, 3, 7, 8, 4, 5}
	var myHeap []*Node
	myHeap = append(myHeap, &Node{})
	for _, value := range arrList {
		tmp := Node{}
		tmp.Value = value
		myHeap = InsertHeap(myHeap, &tmp)
	}
	heapShow(myHeap)
	myHeap = SortHeap(myHeap)
	heapShow(myHeap)
}

func InsertHeap(heaps []*Node, one *Node) []*Node {
	heaps = append(heaps, one)
	length := len(heaps)
	maxPos := length - 1
	//heapShow(heaps)
	heaps = AdjustHeap(heaps, length, maxPos)
	//heapShow(heaps)
	return heaps
}

func SortHeap(heaps []*Node) []*Node {
	length := len(heaps)
	length = length - 1
	if length == 1 {
		return heaps
	}
	if length == 2 {
		heaps = AdjustHeap(heaps, length-1, length)
	}
	for length > 0 {
		SliceNodeSwap(heaps, 1, length)
		length--
		fmt.Println(length)
		Heapfiy(heaps, length, 1)
		heapShow(heaps)
	}
	return heaps
}

func AdjustHeap(heaps []*Node, length int, maxPos int) []*Node {
	if length == 1 {
		return heaps
	}
	if maxPos == 2 {
		if heaps[maxPos].Value > heaps[maxPos-1].Value {
			SliceNodeSwap(heaps, maxPos, maxPos-1)
		}
		return heaps
	}
	//SliceNodeSwap(heaps, 1, maxPos)
	for length/2 > 0 && heaps[maxPos].Value > heaps[length/2].Value {
		SliceNodeSwap(heaps, maxPos, length/2)
		length = length / 2
		maxPos = length
	}
	//heapShow(heaps)
	return heaps
}

func heapShow(heaps []*Node) {
	for one, value := range heaps {
		fmt.Println(one, value)
	}
}

func SliceNodeSwap(list []*Node, i int, j int) {
	x := list[i]
	list[i] = list[j]
	list[j] = x
}

func Heapfiy(heaps []*Node, length int, pos int) {
	for {
		maxPos := pos
		if pos*2 < length && heaps[pos].Value < heaps[pos*2].Value {
			maxPos = pos * 2
		}
		if pos*2+1 < length && heaps[maxPos].Value < heaps[pos*2+1].Value {
			maxPos = pos*2 + 1
		}
		if maxPos == pos {
			break
		}
		SliceNodeSwap(heaps, pos, maxPos)
		pos = maxPos
	}
}
