package main

import "fmt"

/*
地上有一个m行和n列的方格。一个机器人从坐标0,0的格子开始移动，每一次只能向左，右，上，下四个方向移动一格，
但是不能进入行坐标和列坐标的数位之和大于k的格子。 例如，当k为18时，机器人能够进入方格（35,37），因为3+5+3+7 = 18。
但是，它不能进入方格（35,38），因为3+5+3+8 = 19。请问该机器人能够达到多少个格子？
*/

func movingCount(rows, cols, k int) int {
	var rs int
	//记录位置是否走过
	var walked = make([]bool, rows*cols)
	var walk func(i, j int)

	walk = func(i, j int) {
		// 递归结束条件 到达边界
		if i < 0 || j < 0 || i >= rows || j >= cols {
			return
		}

		// 计算走到了哪一个坐标
		index := i*cols + j

		// 如果当前位置没有走过那么才进行处理
		if !walked[index] {
			// 判断当前的位置是否满足要求
			if parseSum(i, j) <= k {
				rs++
			}
			walked[index] = true

			//在继续往4个方向走下去
			walk(i-1, j)
			walk(i+1, j)
			walk(i, j-1)
			walk(i, j+1)
		}
	}

	walk(0, 0)
	return rs
}

func parseSum(i, j int) int {
	var rs int

	for i > 0 {
		rs += i % 10
		i /= 10
	}
	for j > 0 {
		rs += j % 10
		j /= 10
	}

	return rs
}

func main() {
	fmt.Println(movingCount(6, 6, 6))
}
