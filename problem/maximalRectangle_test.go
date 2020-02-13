package problem

import "fmt"

/**
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
问题分析：
1.首先，矩形面积是宽*高
2.其次，求最大面积，所以一般是用动态规划求局部最优解
3.如果从矩形左上角开始最优解一定包含子矩形 所以要从左到右从上倒下的顺序进行遍历
4.局部最优解是要最宽和最高两种情况下，哪一个最大
*/
func maximalRectangle(matrix [][]byte) int {
	v := len(matrix)
	if v <= 0 {
		return 0
	}
	h := len(matrix[0])
	maxArea := 0
	for i := 0; i < v; i++ {
		for j := 0; j < h; j++ {
			if matrix[i][j] == '1' {
				tmpArea := calculateArea(matrix, i, j)
				if tmpArea > maxArea {
					maxArea = tmpArea
				}
			}
		}
	}
	return maxArea
}

func calculateArea(matrix [][]byte, i int, j int) int {
	v := len(matrix)
	h := len(matrix[0])
	width, height := 0, 0
	minWith := h - j
	maxArea := 0
	for m := i; m < v; m++ {
		if matrix[m][j] == '1' {
			height++
		} else {
			break
		}
		width = 0
		for n := j; n < h; n++ {
			if matrix[m][n] == '1' {
				width++
			} else {
				break
			}
		}
		if width < minWith {
			fmt.Println(minWith, height)
			if maxArea < minWith*(height-1) {
				maxArea = minWith * (height - 1)
			}
			minWith = width
		}
	}
	if maxArea < minWith*height {
		maxArea = minWith * height
	}

	return maxArea
}
