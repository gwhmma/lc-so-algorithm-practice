package main

import "fmt"

/*
统计一个数字在排序数组中出现的次数
*/

func countOfK(a []int, k int) int {
	rs := 0

	for _, v := range a {
		if v == k {
			rs++
		} else if rs > 0 {
			break
		}
	}
	return rs
}

func main() {
	a := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7, 8}
	fmt.Println(countOfK(a, 4))
}
