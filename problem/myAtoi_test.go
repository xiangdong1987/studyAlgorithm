package problem

import (
	"fmt"
	"strconv"
	"testing"
)

func myAtoi(str string) int {
	result := 0
	symbol := ""
	start := 0
	for i := 0; i < len(str); i++ {
		if result > INT32_MAX && symbol != "-" {
			return INT32_MAX
		}
		if result > INT32_MAX && symbol == "-" {
			return INT32_MIN
		}
		if (string(str[i]) == "-" || string(str[i]) == "+") && result == 0 {
			fmt.Println(symbol, result, start)
			if symbol != "" {
				break
			}
			if start != 0 {
				break
			}
			symbol = string(str[i])
		} else {
			if (string(str[i]) == " " || string(str[i]) == "0") && result == 0 && symbol == "" {
				if string(str[i]) == "0" {
					start++
				}
				if string(str[i]) == " " && start != 0 {
					break
				}
				continue
			}
			if string(str[i]) >= "0" && string(str[i]) <= "9" {
				x, _ := strconv.Atoi(string(str[i]))
				result = result*10 + x
				start++
			} else {
				break
			}
		}
	}

	if symbol == "-" {
		result = 0 - result
	}
	if result > INT32_MAX && symbol != "-" {
		return INT32_MAX
	}
	if result < INT32_MIN && symbol == "-" {
		return INT32_MIN
	}
	return result
}

func TestMyAtoi(t *testing.T) {
	fmt.Println(myAtoi("0-1"))
}
