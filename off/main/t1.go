package main

import "fmt"

/*
在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func find(s [][]int, target int) bool {
	if len(s) == 0 {
		return false
	}

	min := s[0][0]
	max := s[len(s)-1][len(s[0])-1]

	if target < min || target > max {
		return false
	} else if target == max || target == min {
		return true
	}

	// 取二维数组左下角的数tmp tmp是该行最小的数 是该列最大的数
	// 1.如果target比tmp大那么往该行右移进行比较
	// 2.如果target比tmp小那么往该列上移进行比较
	// 3.对之后的数重复进行 2, 3步骤
	i := len(s) - 1
	j := 0
	tmp := s[i][j]

	for i >= 0 && j < len(s[0]) {
		if target > tmp {
			j++
			if j < len(s[0])-1 {
				tmp = s[i][j]
			}
		} else if target < tmp {
			i--
			if i > 0 {
				tmp = s[i][j]
			}
		} else if target == tmp {
			return true
		}
	}
	return false
}

func main() {
	s := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
		{7, 9, 16, 17},
	}

	fmt.Println(find(s, 4))
}
