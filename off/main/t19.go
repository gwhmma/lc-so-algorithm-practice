package main

import "fmt"

/*
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字，
例如，如果输入如下4 X 4矩阵： 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
  则依次打印出数            1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10
*/

//设置矩阵4个角的边界值
//不断收缩边界值 -1 +1
func printMatrix(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	up := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for {
		//向右走 存入一行
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		//向下走 存入一列
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 向左走 存入一行
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		//向上走 存入一列
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Println(printMatrix(matrix))
}
