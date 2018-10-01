package tool

import (
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
