package problem

import "testing"

func toLowerCase(str string) string {
	result := make([]uint8, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= 65 && str[i] < 97 {
			result[i] = str[i] + 32
		} else {
			result[i] = str[i]
		}
	}
	return string(result)
}

func TestToLowerCase(t *testing.T) {
	println(toLowerCase("Hello"))
}
