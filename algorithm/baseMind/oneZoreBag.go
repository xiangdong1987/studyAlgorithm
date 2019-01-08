package baseMind

import "fmt"

var maxWeight = 0
var itemNum = 5
var bagWeight = 10
var items = []int{4, 5, 2, 7, 1}

func TestOneZero() {
	oneZeroBag(0, 0)
	fmt.Println("maxWeight:", maxWeight)
}
func oneZeroBag(i int, weight int) {
	if weight == bagWeight || i == itemNum {
		fmt.Println(weight, maxWeight)
		if weight > maxWeight {
			maxWeight = weight
		}
		return
	}
	oneZeroBag(i+1, weight)
	if weight+items[i] <= bagWeight {
		oneZeroBag(i+1, weight+items[i])
	}
}
