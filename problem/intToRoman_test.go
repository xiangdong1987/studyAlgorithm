package problem

import (
	"fmt"
	"testing"
)

//应该使用贪心算法 按照优先顺序进行处理
func intToRoman(num int) string {
	result := ""
	i := 0
	for num > 0 {
		tmp := num % 10
		tmpStr := ""
		if tmp == 9 {
			switch i {
			case 0:
				tmpStr += "IX"
				break
			case 1:
				tmpStr += "XC"
				break
			case 2:
				tmpStr += "CM"
				break
			}
			num = num / 10
			i++
			result = tmpStr + result
			continue
		}
		if tmp == 4 {
			switch i {
			case 0:
				tmpStr += "IV"
				break
			case 1:
				tmpStr += "XL"
				break
			case 2:
				tmpStr += "CD"
				break
			}
			num = num / 10
			i++
			result = tmpStr + result
			continue
		}
		if tmp >= 5 {
			switch i {
			case 0:
				tmpStr += "V"
				for x := 0; x < tmp-5; x++ {
					tmpStr += "I"
				}
				break
			case 1:
				tmpStr += "L"
				for x := 0; x < tmp-5; x++ {
					tmpStr += "X"
				}
				break
			case 2:
				tmpStr += "D"
				for x := 0; x < tmp-5; x++ {
					tmpStr += "C"
				}
				break
			}
		}
		if tmp < 5 {
			switch i {
			case 0:
				for x := 0; x < tmp; x++ {
					tmpStr += "I"
				}
				break
			case 1:

				for x := 0; x < tmp; x++ {
					tmpStr += "X"
				}
				break
			case 2:
				for x := 0; x < tmp; x++ {
					tmpStr += "C"
				}
				break
			case 3:
				for x := 0; x < tmp; x++ {
					tmpStr += "M"
				}
				break
			}
		}
		num = num / 10
		i++
		result = tmpStr + result
	}
	return result
}

func TestIntToRoman(t *testing.T) {
	n := intToRoman(10)
	fmt.Println(n)
}
