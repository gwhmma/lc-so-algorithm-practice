package main

import "fmt"

/*
输入一个递增排序的数组和一个数字S，在数组中查找两个数，使得他们的和正好是S，
如果有多对数字的和等于S，输出两个数的乘积最小的
*/

// 首先得到 排序数组的最大值max和最小值min
// 如果 s > max 则最小的2个数为 max 和 s - max
// 如果 s < max 则最小的2个数为 min 和 s - min
func findNum(a []int, s int) (int, int) {
	min := a[0]
	max := a[len(a)-1]

	if s > max {
		return s - max, max
	} else {
		return min, s - min
	}
}

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10,11}
	fmt.Println(findNum(a, 8))
	fmt.Println(findNum(a, 16))
	fmt.Println(findNum(a, 5))
}
