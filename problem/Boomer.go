package problem

//回旋镖问题
func NumberOfBoomerangs(points [][]int) int {
	var sum int
	for i, point1 := range points {
		for j, point2 := range points {
			if i == j {
				continue
			}
			for k, point3 := range points {
				if k == j {
					continue
				}
				if (mySquare(point1[0]-point2[0]) + mySquare(point1[1]-point2[1])) == (mySquare(point1[0]-point3[0]) + mySquare(point1[1]-point3[1])) {
					sum++
				} else {
					continue
				}
			}
		}
	}
	return sum
}
func mySquare(x int) int {
	x = x * x
	return x
}
