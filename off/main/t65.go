package main

import "fmt"

/*
请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
路径可以从矩阵中的任意一个格子开始，每一步可以在矩阵中向左，向右，向上，向下移动一个格子。
如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
例如 a b c e s f c s a d e e 矩阵中包含一条字符串"bcced"的路径，但是矩阵中不包含"abcb"路径，
因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入该格子
*/

// 回溯法
// 用一个相同大小的数组标记相应的位置是否已经走过
// 递归的寻找下一个满足条件的位置
// 递归结束条件: 1 到达数组边界 2 不满足相等条件  3 找到了路径
func hasPath(matrix []byte, rows, cols int, str []byte) bool {
	// 标记数组 用来标记某个位置是否已经走过
	var flag = make([]bool, rows*cols)
	var path func([]byte, int, int, int, int, []bool, []byte, int) bool

	path = func(matrix []byte, i int, j int, rows int, cols int, flag []bool, str []byte, k int) bool {
		// 计算matrix和flag里元素的位置   因为matrix用一维数组表示
		idx := i*cols + j

		//递归结束条件
		//到达边界 或者当前位置元素不匹配或者该位置已经走过了
		if i < 0 || j < 0 || i >= rows || j >= cols || matrix[idx] != str[k] || flag[idx] {
			return false
		}

		// 已经到达str的最后一位说明匹配成功
		if k == len(str)-1 {
			return true
		}

		// 标记这一个位置已经走过
		flag[idx] = true

		// 从现在的位置继续往4个方向走下去  只要有一条路径匹配就返回true
		if path(matrix, i+1, j, rows, cols, flag, str, k+1) ||
			path(matrix, i, j+1, rows, cols, flag, str, k+1) ||
			path(matrix, i-1, j, rows, cols, flag, str, k+1) ||
			path(matrix, i, j-1, rows, cols, flag, str, k+1) {
			return true
		}

		// 如果该位置没有匹配成功那么回溯 将该位置标记改为false
		flag[idx] = false
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if path(matrix, i, j, rows, cols, flag, str, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	s1 := "abcesfcsadee"
	s2 := "bccedf"
	matrix := []byte(s1)
	str := []byte(s2)

	fmt.Println(hasPath(matrix, 3, 4, str))
}
