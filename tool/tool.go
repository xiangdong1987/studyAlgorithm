package tool

import (
	"fmt"
	"math/rand"
	"time"
)

//生成范围随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}

//生成随机数
func GenerateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	//fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func SliceSwap(list []int, i int, j int) {
	x := list[i]
	list[i] = list[j]
	list[j] = x
}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Power(x int, n int) int {
	ans := 1

	for n != 0 {
		ans *= x
		n--
	}
	return ans
}
