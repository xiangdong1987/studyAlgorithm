package problem

import (
	"fmt"
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"
	t1, err := time.Parse(layout, date1)
	if err != nil {
		fmt.Println(err)
	}
	t2, err := time.Parse(layout, date2)
	if err != nil {
		fmt.Println(err)
	}
	res := int(math.Abs(float64((t2.Unix() - t1.Unix()) / (3600 * 24))))
	return res
}
